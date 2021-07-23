package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mxAge18/golang-gin-poc/entity"
	"github.com/mxAge18/golang-gin-poc/service"
	"github.com/mxAge18/golang-gin-poc/validators"
)

type VideoController interface{
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	 service service.VideoService
}

var validate *validator.Validate
func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service:service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video) 
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}