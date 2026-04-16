package interactor

import (
	"go_practice/domain/entity"
	inputport "go_practice/usecase/port/input"
	outputport "go_practice/usecase/port/output"
)

type UserUseCase struct {
	userRepository outputport.UserRepository
}

func NewUserUseCase(userRepo outputport.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepo,
	}
}

func (uc *UserUseCase) FindByID(userID string) (entity.User, error) {
	return uc.userRepository.FindByID(userID)
}

func (uc *UserUseCase) List(query inputport.ListUsersQuery) ([]entity.User, int, *int, error) {
	users, total, nextPage, err := uc.userRepository.Search(outputport.UserSearch{
		Page:     query.Page,
		Take:     query.Take,
		Q:        query.Q,
		Order:    query.Order,
		OrderBy:  query.OrderBy,
		UserType: query.UserType,
	})
	if err != nil {
		return nil, 0, nil, err
	}

	return users, total, nextPage, nil
}
