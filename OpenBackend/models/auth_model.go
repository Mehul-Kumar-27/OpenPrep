package models


type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}


type AuthResponse struct {
	Message string
	Status  int
	Flag    bool
}