package repository

import (
	"errors"
	"fmt"
	"go_practice/adapter/database/model"
	"go_practice/domain/entity"
	"go_practice/usecase/interactor"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(userID string) (_ entity.User, err error) {
	var user model.User
	err = r.db.Where("id = ?", userID).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, fmt.Errorf("%w: user not found",interactor.ErrKind.NotFound)
	}
	if err != nil {
		return entity.User{}, fmt.Errorf("%w: failed to find user by ID: %v", interactor.ErrKind.DB, err)
	}

	return user.ToEntity(), nil
}

func (r *UserRepository) FindByEmail(email string) (_ entity.User, err error) {
	var user model.User
	err = r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, fmt.Errorf("%w: user not found",interactor.ErrKind.NotFound)
	}
	if err != nil {
		return entity.User{}, fmt.Errorf("%w: failed to find user by email: %v", interactor.ErrKind.DB, err)
	}

	return user.ToEntity(), nil
}

func (r *UserRepository) Create(user entity.User) error {
	if err := r.db.Create(&model.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     string(user.Role),
	}).Error; err != nil {
		return fmt.Errorf("%w: failed to create user: %v", interactor.ErrKind.DB, err)
	}
	return nil
}