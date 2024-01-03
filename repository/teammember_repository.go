package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITeamMemberRepository interface {
	// チームに参加
	AssignToTeam(teamMember *model.TeamMember) error
	// チームから外れる
	UnassignFromTeam(teamMember *model.TeamMember, userId uint, teamId uint) error
	// ユーザーIDからチームを取得
	GetTeamMembersByTeamId(teamMember *[]model.TeamMember, userId uint) error
}

type teamMemberRepository struct {
	db *gorm.DB
}

func NewTeamMemberRepository(db *gorm.DB) ITeamMemberRepository {
	return &teamMemberRepository{db}
}

func (tmr *teamMemberRepository) AssignToTeam(teamMember *model.TeamMember) error {
	if err := tmr.db.Create(teamMember).Error; err != nil {
		return err
	}

	return nil
}

func (tmr *teamMemberRepository) UnassignFromTeam(teamMember *model.TeamMember, userId uint, teamId uint) error {
	result := tmr.db.Model(teamMember).Clauses(clause.Returning{}).Where("user_id=? AND team_id=?", userId, teamId).Update("delete_flg", true)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("team member does not exist")
	}

	return nil
}

func (tmr *teamMemberRepository) GetTeamMembersByTeamId(teamMembers *[]model.TeamMember, userId uint) error {
	result := tmr.db.Where("userId=?", userId).Find(teamMembers);
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}