package controller

import (
	"goapp-gin/dto"
	"goapp-gin/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func NewLoginController(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (con *loginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := con.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return con.jWtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
