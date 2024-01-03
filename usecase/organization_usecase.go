package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IOrganizationUseCase interface {
	// 作成した組織を表示する
	GetCreatedOrganizationsByUserId(userId uint) ([]model.OrganizationResponse, error)
	// 組織を作成する
	CreateOrganization(organization model.Organization) (model.OrganizationResponse, error)
	// 組織の一覧を取得する
	ListOrganizations() ([]model.OrganizationResponse, error)
}

type organizationUseCase struct {
	or repository.IOrganizationRepository
}

func NewOrganizationUseCase(or repository.IOrganizationRepository) IOrganizationUseCase {
	return &organizationUseCase{or}
}

func (ou *organizationUseCase) GetCreatedOrganizationsByUserId(userId uint) ([]model.OrganizationResponse, error) {
	organizations := make([]model.Organization, 0)
	if err := ou.or.GetCreatedOrganizationsByUserId(&organizations, userId); err != nil {
		return nil, err
	}

	ResOrganizations := make([]model.OrganizationResponse, len(organizations))
	for i, v := range organizations {
		ResOrganizations[i] = model.OrganizationResponse {
			ID: v.ID,
			Name: v.Name,
			Description: v.Description,
		}
	}

	return ResOrganizations, nil
}

func (ou *organizationUseCase) CreateOrganization(organization model.Organization) (model.OrganizationResponse, error) {
	if err := ou.or.CreateOrganization(&organization); err != nil {
		return model.OrganizationResponse{}, err
	}
	resOrganization := model.OrganizationResponse {
		ID: organization.ID,
		Name: organization.Name,
		Description: organization.Description,
	}

	return resOrganization, nil
}

func (ou *organizationUseCase) ListOrganizations() ([]model.OrganizationResponse, error) {
	organizations := make([]model.Organization, 0)
	if err := ou.or.ListOrganizations(&organizations); err != nil {
		return nil, err
	}
	resOrganizations := make([]model.OrganizationResponse, len(organizations))
	for i, v := range organizations {
		resOrganizations[i] = model.OrganizationResponse{
			ID: v.ID,
			Name: v.Name,
			Description: v.Description,
		}
	}

	return resOrganizations, nil
}