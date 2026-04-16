package input_port

import "go_practice/domain/entity"

type SignupInput struct {
	Email    string
	Password string
	Name     string
}

type IAuthUseCase interface {
	Signup(input SignupInput) (entity.User, error)
	Login(email, password string) (entity.User, string, error)
}
