package input

import "go_practice/domain/entity"

type ListUsersQuery struct {
	Page     int
	Take     int
	Q        string
	Order    string
	OrderBy  string
	UserType string
}

type UpdateByMeInput struct {
	ID   string
	Name string
}

type IUserUseCase interface {
	FindByID(userID string) (entity.User, error)
	List(query ListUsersQuery) ([]entity.User, int, *int, error)
	UpdateByMe(input UpdateByMeInput) (entity.User, error)
}
