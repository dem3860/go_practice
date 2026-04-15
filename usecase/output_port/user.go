package output_port

import "go_practice/domain/entity"

type UserRepository interface {
	Create(entity.User) error
	FindByID(userID string) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}
