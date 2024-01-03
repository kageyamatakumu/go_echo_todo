package model

import "time"

type Task struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Title     string     `json:"title" gorm:"not null"`
	Status    TaskStatus `json:"status" gorm:"not null; default:0"`
	Memo      string     `json:"memo" gorm:"size: 65535"`
	DeadLine  time.Time  `json:"dead_line" gorm:"not null; default:CURRENT_TIMESTAMP; type:date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Team			Team       `json:"team" gorm:"foreignKey:TeamId; constraint:OnDelete:CASCADE"`
	TeamId    uint       `json:"team_id" gorm:"not null"`
}

type TaskResponse struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Title     string     `json:"title" gorm:"not null"`
	Status    TaskStatus `json:"status" gorm:"not null; default:0"`
	Memo      string     `json:"memo" gorm:"size: 65535"`
	DeadLine  time.Time  `json:"dead_line" gorm:"not null; default:CURRENT_TIMESTAMP; type:date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type TaskStatus int

const (
	TaskStatusUnstarted TaskStatus = iota
	TaskStatusStarted
	TaskStatusCompleted
)

