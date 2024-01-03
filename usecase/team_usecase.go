package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type ITeamUseCase interface {
	// 所属チームを取得する
	GetAssignTeamByUserId(userId uint) ([]model.TeamResponse, error)
	// チームを作成する
	CreateTeam(team model.Team) (model.TeamResponse, error)
	// 組織内のチーム一覧を取得する
	GetTeamsByOrganizationId(organizationId uint) ([]model.TeamResponse, error)
	// チームを削除する
	DeleteTeam(teamId uint, userId uint) error
}

type teamUseCase struct {
	tr repository.ITeamRepository
	tmr repository.ITeamMemberRepository
}

func NewTeamUseCase(tr repository.ITeamRepository, tmr repository.ITeamMemberRepository) ITeamUseCase {
	return &teamUseCase{tr, tmr}
}

func (tu *teamUseCase) GetAssignTeamByUserId(userId uint) ([]model.TeamResponse, error) {
	teams := make([]model.Team, 0)
	if err := tu.tr.GetAssignTeamByUserId(&teams, userId); err != nil {
		return []model.TeamResponse{}, err
	}

	teamsRes := make([]model.TeamResponse, len(teams))
	for i, v := range teams {
		teamsRes[i] = model.TeamResponse{
			ID: v.ID,
			Name: v.Name,
			Description: v.Description,
		}
	}

	return teamsRes, nil
}

func (tu *teamUseCase) CreateTeam(team model.Team) (model.TeamResponse, error) {
	if err := tu.tr.CreateTeam(&team); err != nil {
		return model.TeamResponse{}, nil
	}

	resTeam := model.TeamResponse {
		ID: team.ID,
		Name: team.Name,
		Description: team.Description,
	}

	return resTeam, nil
}

func (tu *teamUseCase) GetTeamsByOrganizationId(organizationId uint) ([]model.TeamResponse, error) {
	teams := make([]model.Team, 0)
	if err := tu.tr.GetTeamsByOrganizationId(&teams, organizationId); err != nil {
		return []model.TeamResponse{}, err
	}

	resTeams := make([]model.TeamResponse, len(teams))
	for i, v := range teams {
		resTeams[i] = model.TeamResponse{
			ID: v.ID,
			Name: v.Name,
			Description: v.Description,
		}
	}

	return resTeams, nil
}

func (tu *teamUseCase) DeleteTeam(teamId uint, userId uint) error {
	teamMember := model.TeamMember{}
	if err := tu.tmr.UnassignFromTeam(&teamMember, userId, teamId); err != nil {
		return err
	}
	if err := tu.tr.DeleteTeam(teamId); err != nil {
		return err
	}

	return nil
}

