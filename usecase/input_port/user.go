package input_port

import "go_practice/domain/entity"

type UserCreate struct {
	Email    string
	Password string
	Name     string
}

type IUserUseCase interface {
	Create(input UserCreate) (entity.User,error)
}
