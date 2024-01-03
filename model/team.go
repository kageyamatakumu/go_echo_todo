package model

type Team struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	Name           string       `json:"name" gorm:"not null"`
	Description    string       `json:"description" gorm:"size: 65535"`
	Organization   Organization `json:"organization" gorm:"foreignKey:OrganizationId; constraint:OnDelete:CASCADE"`
	OrganizationId uint         `json:"organization_id" gorm:"not null"`
}

type TeamResponse struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	Name           string       `json:"name" gorm:"not null"`
	Description    string       `json:"description" gorm:"size: 65535"`
}