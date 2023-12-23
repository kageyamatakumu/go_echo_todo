package model

import "time"

type Task struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Title     string     `json:"title" gorm:"not null"`
	Status    TaskStatus `json:"status" gorm:"not null; default:0"`
	Memo      string     `json:"memo" gorm:"size: 65535"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	User      User       `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint       `json:"userId" gorm:"not null"`
}

type TaskResponse struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Title     string     `json:"title" gorm:"not null"`
	Status    TaskStatus `json:"status" gorm:"not null; default:0"`
	Memo      string     `json:"memo" gorm:"size: 65535"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type TaskStatus int

const (
	TaskStatusUnstarted TaskStatus = iota
	TaskStatusStarted
	TaskStatusCompleted
)

