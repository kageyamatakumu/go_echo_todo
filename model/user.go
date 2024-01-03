package model

import "time"

type User struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	Email          string       `json:"email" gorm:"unique"`
	Password       string       `json:"password"`
	Name           string       `json:"name"`
	Organization   Organization `json:"organization" gorm:"foreignKey:OrganizationId; constraint:OnDelete:CASCADE"`
	OrganizationId uint         `json:"organization_id" gorm:"default:1"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
	Name  string `json:"name"`
}

type UserAssignResponse struct {
	OrganizationId uint         `json:"organization_id"`
}