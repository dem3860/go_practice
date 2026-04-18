package model

import (
	"go_practice/domain/entity"
	"time"
)

type User struct {
	ID        string    `gorm:"type:varchar(26);primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(50);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	Role      string    `gorm:"type:varchar(20);not null;default:'user'" json:"role"`
	CreatedAt time.Time `gorm:"column:created;not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:modified;not null" json:"updatedAt"`
}

func (User) TableName() string {
	return "users"
}

func (u User) ToEntity() entity.User {
	return entity.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Role:      entity.UserRole(u.Role),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
