package repository

import "github.com/alefiengo/go-jwt-app/model"

type UsersRepository interface {
	Save(users model.Users)
	Update(users model.Users)
	Delete(usersId int)
	FindById(usersId int) (model.Users, error)
	FindAll() []model.Users
	FindByUsername(username string) (model.Users, error)
}
