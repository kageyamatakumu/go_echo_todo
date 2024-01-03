package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
)

type ITeamRepository interface {
	// 所属チームを取得する
	GetAssignTeamByUserId(teams *[]model.Team, userId uint) error
	// チームを作成する
	CreateTeam(team *model.Team) error
	// 組織内のチーム一覧を取得する
	GetTeamsByOrganizationId(teams *[]model.Team, organizationId uint) error
	// チームを削除する
	DeleteTeam(teamId uint) error
}

type teamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) ITeamRepository {
	return &teamRepository{db}
}

func (tr *teamRepository) GetAssignTeamByUserId(teams *[]model.Team, userId uint) error {
	if err := tr.db.
	Table("teams").
	Select("teams.*").
	Joins("INNER JOIN team_members ON team_members.team_id = teams.id").
	Where("team_members.user_id = ?", userId).
	Find(teams).Error; err != nil {
		return err
	}

	return nil
}

func (tr *teamRepository) CreateTeam(team *model.Team) error {
	if err := tr.db.Create(team).Error; err != nil {
		return err
	}

	return nil
}

func (tr *teamRepository) GetTeamsByOrganizationId(teams *[]model.Team, organizationId uint) error {
	if err := tr.db.Joins("Organization").Where("organization_id", organizationId).Order("team_id").Find(teams).Error; err != nil {
		return err
	}

	return nil
}

func (tr *teamRepository) DeleteTeam(teamId uint) error {
	result := tr.db.Where("id=?", teamId).Delete(model.Team{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

