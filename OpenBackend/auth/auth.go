package auth


type AuthService interface {
	ManualAuthService
}

type AuthController struct {
	AuthService
	manualAuthService ManualAuthService
}


func NewAuthController() *AuthController {
	manualAuthController := NewManualAuthController()

	return &AuthController{
		manualAuthService: manualAuthController,
	}
}