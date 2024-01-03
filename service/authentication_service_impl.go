package service

import (
	"errors"

	"github.com/alefiengo/go-jwt-app/config"
	"github.com/alefiengo/go-jwt-app/data/request"
	"github.com/alefiengo/go-jwt-app/helper"
	"github.com/alefiengo/go-jwt-app/model"
	"github.com/alefiengo/go-jwt-app/repository"
	"github.com/alefiengo/go-jwt-app/utils"

	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewAuthenticationServiceImpl(usersRepository repository.UsersRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepository: usersRepository,
		Validate:        validate,
	}
}

// Login implements AuthenticationService.
func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {
	// Find username in database
	new_users, users_err := a.UsersRepository.FindByUsername(users.Username)

	if users_err != nil {
		return "", errors.New("Invalid username or password")
	}

	config, _ := config.LoadConfig(".")
	verify_error := utils.VerifyPassword(new_users.Password, users.Password)

	if verify_error != nil {
		return "", errors.New("Invalid username or password")
	}

	// Generate token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_users.Id, config.TokenSecret)
	helper.ErrorPanic(err_token)

	return token, nil
}

// Register implements AuthenticationService.
func (a *AuthenticationServiceImpl) Register(users request.CreateUserRequest) {
	hashedPassword, err := utils.HashPassword(users.Password)
	helper.ErrorPanic(err)

	newUser := model.Users{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
	}

	a.UsersRepository.Save(newUser)
}
