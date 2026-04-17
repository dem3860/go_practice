package factory

import (
	"go_practice/domain/entity"
	"go_practice/domain/validation"
)

type NewUserArgs struct {
	ID       string
	Name     string
	Role     string
	Email    string
	Password string
}

func NewUser(arg NewUserArgs) (entity.User, error) {
	err := validation.ValidateName(arg.Name)
	if err != nil {
		return entity.User{}, err
	}

	err = validation.ValidateEmail(arg.Email, false)
	if err != nil {
		return entity.User{}, err
	}
	if err = validation.ValidatePassword(arg.Password); err != nil {
		return entity.User{}, err
	}

	return entity.User{
		ID:       arg.ID,
		Name:     arg.Name,
		Role:     entity.UserRole(arg.Role),
		Email:    arg.Email,
		Password: arg.Password,
	}, nil
}
