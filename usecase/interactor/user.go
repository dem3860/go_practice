package interactor

import "go_practice/domain/entity"
import "go_practice/usecase/port/output"

type UserUseCase struct {
	userRepository output.UserRepository
}

func NewUserUseCase(userRepo output.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepo,
	}
}

func (uc *UserUseCase) FindByID(userID string) (entity.User, error) {
	return uc.userRepository.FindByID(userID)
}
