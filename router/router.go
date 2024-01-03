package router

import (
	"net/http"

	"github.com/alefiengo/go-jwt-app/controller"
	"github.com/alefiengo/go-jwt-app/middleware"
	"github.com/alefiengo/go-jwt-app/repository"

	"github.com/gin-gonic/gin"
)

func NewRouter(userRepository repository.UsersRepository, authenticationController *controller.AuthenticationController, userController *controller.UserController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home")
	})

	router := service.Group("/api/v1")

	authenticationRouter := router.Group("/auth")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.Login)

	usersRouter := router.Group("/users")
	usersRouter.GET("", middleware.DeserializeUser(userRepository), userController.GetUsers)

	return service
}
