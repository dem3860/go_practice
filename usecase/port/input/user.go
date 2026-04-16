package input

import "go_practice/domain/entity"

type IUserUseCase interface {
	FindByID(userID string) (entity.User, error)
}
