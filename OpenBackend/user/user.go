package user

import (
	"github.com/Mehul-Kumar-27/OpenBackend/database"
	"gorm.io/gorm"
)

type UserService interface {
	EmailExists(email string) bool
}

type UserController struct {
	userService UserService
	db          *gorm.DB
}

func NewUserController() *UserController {
	return &UserController{
		db : database.PostgresConnHandler.DB,
	}
}


func (controller *UserController) EmailExists(email string) bool {
	return false
}
