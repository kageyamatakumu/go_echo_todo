package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IOrganizationRepository interface {
	// 作成した組織を表示する
	GetCreatedOrganizationsByUserId(organizations *[]model.Organization, userId uint) error
	// 組織を作成する
	CreateOrganization(organization *model.Organization) error
	// 組織の一覧を取得する
	ListOrganizations(organizations *[]model.Organization) error
}

type organizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) IOrganizationRepository {
	return &organizationRepository{db}
}

// 作成した組織を表示する
func (or *organizationRepository) GetCreatedOrganizationsByUserId(organizations *[]model.Organization, userId uint) error {
	if err := or.db.Where("founder=?", userId).Find(organizations).Error; err != nil {
		return err
	}

	return nil
}

// 組織を作成する
func (or *organizationRepository) CreateOrganization(organization *model.Organization) error {
	if err := or.db.Create(organization).Error; err != nil {
		return err
	}

	return nil
}

// 組織の一覧を取得する
func (or *organizationRepository) ListOrganizations(organizations *[]model.Organization) error {
	if err := or.db.Find(organizations).Error; err != nil {
		return err
	}

	return nil
}

