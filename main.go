package main

import (
	"io"
	_ "net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mxAge18/golang-gin-poc/api"
	"github.com/mxAge18/golang-gin-poc/controller"
	"github.com/mxAge18/golang-gin-poc/docs"
	"github.com/mxAge18/golang-gin-poc/middlewares"
	"github.com/mxAge18/golang-gin-poc/repository"
	"github.com/mxAge18/golang-gin-poc/service"
	_ "github.com/tpkeeper/gin-dump"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
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

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Pragmatic Reviews - Video API"
	docs.SwaggerInfo.Description = "Pragmatic Reviews - Youtube Video API Learn Gin."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:9090"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}


	setupLogOutput()
	server := gin.New()
	// server.Static("/css", "./templates/css")
	// server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger())

	// viewRoutes := server.Group("/view")
	// {
	// 	viewRoutes.GET("/videos", videoController.ShowAll)
	// }

	videoAPI := api.NewVideoAPI(loginController, videoController)
	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.POST("/token", videoAPI.Authenticate)
		}

		videos := apiRoutes.Group("/videos", middlewares.AuthorizeJWT())
		{
			videos.GET("", videoAPI.GetVideos)
			videos.POST("", videoAPI.CreateVideo)
			videos.PUT(":id", videoAPI.UpdateVideo)
			videos.DELETE(":id", videoAPI.DeleteVideo)
		}
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}