package input_port

import "go_practice/domain/entity"

type UserCreate struct {
	Email    string
	Password string
	Name     string
	Role     string
}

type IUserUseCase interface {
	Create(input UserCreate) (entity.User, error)
	Login(email, password string) (entity.User, string, error)
}
