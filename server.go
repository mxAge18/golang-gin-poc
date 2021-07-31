package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mxAge18/golang-gin-poc/controller"
	"github.com/mxAge18/golang-gin-poc/middlewares"
	"github.com/mxAge18/golang-gin-poc/repository"
	"github.com/mxAge18/golang-gin-poc/service"
	ginDump "github.com/tpkeeper/gin-dump"
)

var(
	repoService repository.VideoRepository = repository.NewVideoRepostory()
	videoService service.VideoService = service.New(repoService)
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
		apiRoutes.PUT("/videos/:id", func(c *gin.Context) {
			err := videoController.Update(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error()})
			} else {
				c.JSON(http.StatusAccepted, gin.H{ "message": "video is updated"})
			}
		})
		apiRoutes.DELETE("/videos/:id", func(c *gin.Context) {
			err := videoController.Delete(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error()})
			} else {
				c.JSON(http.StatusAccepted, gin.H{ "message": "video is deleted"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}