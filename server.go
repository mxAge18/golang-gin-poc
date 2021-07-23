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
)
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(),
	middlewares.BasicAuth(), ginDump.Dump())
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{ "error": err.Error()})
		} else {
			ctx.JSON(http.StatusAccepted, gin.H{ "message": "video input is valid"})
		}
	})
	server.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}