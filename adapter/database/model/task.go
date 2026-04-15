package model

import "time"

type Task struct {
	ID        string    `gorm:"type:varchar(26);primaryKey" json:"id"`
	UserID    string    `gorm:"type:varchar(26);not null;index" json:"userId"`
	Title     string    `gorm:"type:varchar(128);not null" json:"title"`
	Status    string    `gorm:"type:varchar(20);not null" json:"status"`
	CreatedAt time.Time `gorm:"column:created;not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:modified;not null" json:"modified"`

	User User `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
}

func (Task) TableName() string {
	return "tasks"
}
