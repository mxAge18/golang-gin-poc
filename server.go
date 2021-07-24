package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mxAge18/golang-gin-poc/controller"
	"github.com/mxAge18/golang-gin-poc/middlewares"
	"github.com/mxAge18/golang-gin-poc/service"
	ginDump "github.com/tpkeeper/gin-dump"
)

var(
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
	loginService service.LoginService = service.NewLoginService()
	jwtService service.JwtService = service.NewJWTService()
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger())
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT(), ginDump.Dump())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{ "error": err.Error()})
			} else {
				ctx.JSON(http.StatusAccepted, gin.H{ "message": "video input is valid"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}