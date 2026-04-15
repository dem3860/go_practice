package constructor

import (
	"go_practice/domain/entity"
	"go_practice/domain/validation"
)

type NewUserCreateArgs struct {
	ID                   string
	Name					 string
	Email                    string
	Password                 string
}

func NewUserCreate(arg NewUserCreateArgs) (entity.User, error) {
	err := validation.ValidateEmail(arg.Email, false)
	if err != nil {
		return entity.User{}, err
	}
	if err = validation.ValidatePassword(arg.Password); err != nil {
		return entity.User{}, err
	}

	return entity.User{
		ID:   arg.ID,
		Name:     arg.Name,
		Email:    arg.Email,
		Password: arg.Password,
	}, nil
}
