package service

import "github.com/alefiengo/go-jwt-app/go-jwt-app/data/request"

type AuthenticationService interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUserRequest)
}
