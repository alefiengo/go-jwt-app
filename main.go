package main

import (
	"log"
	"net/http"

	"github.com/alefiengo/go-jwt-app/go-jwt-app/controller"
	"github.com/alefiengo/go-jwt-app/go-jwt-app/helper"
	"github.com/alefiengo/go-jwt-app/go-jwt-app/model"
	"github.com/alefiengo/go-jwt-app/go-jwt-app/repository"
	"github.com/alefiengo/go-jwt-app/go-jwt-app/router"
	"github.com/alefiengo/go-jwt-app/go-jwt-app/service"

	"github.com/alefiengo/go-jwt-app/go-jwt-app/config"
	"github.com/go-playground/validator/v10"
)

func main() {
	loadConfig, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	// Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})

	// Init Repository
	usersRepository := repository.NewUsersRepositoryImpl(db)

	// Init Service
	authenticationService := service.NewAuthenticationServiceImpl(usersRepository, validate)

	// Init Controller
	authenticationController := controller.NewAuthenticationController(authenticationService)
	usersController := controller.NewUsersController(usersRepository)

	routes := router.NewRouter(usersRepository, authenticationController, usersController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
