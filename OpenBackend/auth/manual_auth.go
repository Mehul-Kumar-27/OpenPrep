package auth

import (
	"github.com/Mehul-Kumar-27/OpenBackend/models"
	"github.com/gin-gonic/gin"
)

type ManualAuthService interface {
	Login(ctx *gin.Context)
	Signup(ctx *gin.Context)
}

type ManualAuthController struct {
	manualAuthService ManualAuthService
}

func NewManualAuthController() ManualAuthService {
	return &ManualAuthController{}
}

func (controller *ManualAuthController) Login(ctx *gin.Context) {

}

func (controller *ManualAuthController) Signup(ctx *gin.Context) {
	//// Get the username , password, email from the request
	var authRequest models.AuthRequest

	if err := ctx.ShouldBindJSON(&authRequest); err != nil {
		var authResponse models.AuthResponse
		authResponse.Message = "Invalid request"
		authResponse.Status = 400
		authResponse.Flag = false

		ctx.JSON(400, authResponse)
	}

	// Check if the user already exists
	

}
