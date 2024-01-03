package controller

import (
	"net/http"

	"github.com/alefiengo/go-jwt-app/go-jwt-app/data/response"
	"github.com/alefiengo/go-jwt-app/go-jwt-app/repository"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository repository.UsersRepository
}

func NewUsersController(repository repository.UsersRepository) *UserController {
	return &UserController{userRepository: repository}
}

func (controller *UserController) GetUsers(ctx *gin.Context) {
	users := controller.userRepository.FindAll()
	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully fetch all user data!",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
