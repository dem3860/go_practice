package repository

import (
	"errors"
	"fmt"
	"go_practice/adapter/database/model"
	"go_practice/domain/entity"
	"go_practice/usecase/interactor"
	outputport "go_practice/usecase/port/output"
	"strings"

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
		return entity.User{}, fmt.Errorf("%w: user not found", interactor.ErrKind.NotFound)
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
		return entity.User{}, fmt.Errorf("%w: user not found", interactor.ErrKind.NotFound)
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

func (r *UserRepository) Update(user entity.User) error {
	if err := r.db.Model(&model.User{}).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"name":  user.Name,
			"email": user.Email,
			"role":  string(user.Role),
		}).Error; err != nil {
		return fmt.Errorf("%w: failed to update user: %v", interactor.ErrKind.DB, err)
	}
	return nil
}

func (r *UserRepository) Search(query outputport.UserSearch) (_ []entity.User, total int, nextPage *int, err error) {
	db := r.db.Model(&model.User{})

	if query.Q != "" {
		keyword := "%" + strings.ToLower(query.Q) + "%"
		db = db.Where("LOWER(name) LIKE ? OR LOWER(email) LIKE ?", keyword, keyword)
	}
	if query.UserType != "" {
		db = db.Where("role = ?", query.UserType)
	}

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, nil, fmt.Errorf("%w: failed to count users: %v", interactor.ErrKind.DB, err)
	}
	total = int(count)

	orderByMap := map[string]string{
		"createdAt": "created",
		"name":      "name",
		"email":     "email",
		"role":      "role",
	}

	offset := (query.Page - 1) * query.Take
	var users []model.User
	if err := db.Order(orderByMap[query.OrderBy] + " " + query.Order).
		Limit(query.Take).
		Offset(offset).
		Find(&users).Error; err != nil {
		return nil, 0, nil, fmt.Errorf("%w: failed to search users: %v", interactor.ErrKind.DB, err)
	}

	result := make([]entity.User, 0, len(users))
	for _, user := range users {
		result = append(result, user.ToEntity())
	}

	if query.Page*query.Take < total {
		page := query.Page + 1
		nextPage = &page
	}

	return result, total, nextPage, nil
}

func (r *UserRepository) Delete(userD string) (err error) {
	if err := r.db.Model(&model.User{}).
		Where("id = ?", userD).
		Delete(nil).Error; err != nil {
		return fmt.Errorf("%w: failed to delete user: %v", interactor.ErrKind.DB, err)
	}
	return nil
}
