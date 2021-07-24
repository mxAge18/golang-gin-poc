package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mxAge18/golang-gin-poc/dto"
	"github.com/mxAge18/golang-gin-poc/service"
)

type LoginController interface{
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService service.JwtService
}

func NewLoginController(loginService service.LoginService,
	jwtService service.JwtService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (service *loginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials) 
	if err != nil {
		return ""
	}
	isAuthed := service.loginService.Login(credentials.Username, credentials.Password)
	if isAuthed {
		return service.jwtService.GenerateToken(credentials.Username, true)
	}
	return ""
}

