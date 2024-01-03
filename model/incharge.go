package model

type InCharge struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	TaskID uint `json:"task_id"`
  UserID uint `json:"user_id"`
}