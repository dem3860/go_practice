package output

import "go_practice/domain/entity"

type UserSearch struct {
	Page     int
	Take     int
	Q        string
	Order    string
	OrderBy  string
	UserType string
}

type UserRepository interface {
	Create(entity.User) error
	Update(entity.User) error
	FindByID(userID string) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	Search(query UserSearch) ([]entity.User, int, *int, error)
}
