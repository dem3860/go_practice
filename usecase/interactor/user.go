package interactor

import (
	"errors"
	"fmt"
	"go_practice/domain/entity"
	"go_practice/domain/validation"
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

func (uc *UserUseCase) UpdateByMe(input inputport.UpdateByMeInput) (entity.User, error) {
	user, err := uc.userRepository.FindByID(input.ID)
	if err != nil {
		if errors.Is(err, ErrKind.NotFound) {
			return entity.User{}, fmt.Errorf("%w: user not found", ErrKind.NotFound)
		}
		return entity.User{}, err
	}

	if err := validation.ValidateName(input.Name); err != nil {
		return entity.User{}, fmt.Errorf("%w: %v", ErrKind.Validation, err)
	}

	user.Name = input.Name

	if err := uc.userRepository.Update(user); err != nil {
		return entity.User{}, err
	}

	updatedUser, err := uc.userRepository.FindByID(user.ID)
	if err != nil {
		return entity.User{}, err
	}

	return updatedUser, err
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

func (uc *UserUseCase) Delete(userID string) error {
	user, err := uc.userRepository.FindByID(userID)
	if err != nil {
		if errors.Is(err, ErrKind.NotFound) {
			return fmt.Errorf("%w: user not found", ErrKind.NotFound)
		}
		return err
	}
	if err := uc.userRepository.Delete(user.ID); err != nil {
		return err
	}
	return nil
}
