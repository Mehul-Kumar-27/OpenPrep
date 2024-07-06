package auth

import "github.com/gin-gonic/gin"

type LoginInterface interface {
	Login(ctx *gin.Context) 
}

type LoginController struct {
}

func NewLoginController() LoginInterface {
	return &LoginController{}
}

func (controller *LoginController) Login(ctx *gin.Context) {

}
