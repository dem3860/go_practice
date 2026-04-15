package entity

import "time"

type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusDone       TaskStatus = "done"
)

type Task struct {
	ID        string
	UserID    string
	Title     string
	Status    TaskStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}