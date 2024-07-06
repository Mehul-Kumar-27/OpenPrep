package auth

type AuthResponse struct {
	Message string
	Status  int
	Flag    bool
}

type AuthService interface {
	LoginInterface
	
}

type AuthController struct {
	AuthService
	login LoginInterface
}


func NewAuthController() *AuthController {
	login := NewLoginController()

	return &AuthController{
		login: login,
	}
}
