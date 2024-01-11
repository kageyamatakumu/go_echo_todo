package usecase

import (
	"fmt"
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v4"
)

type IUserUseCase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
	// ログインしているユーザーの情報を取得
	GetLoggedInUserDetails(user model.User, userId uint) (model.UserResponse, error)
	// ユーザーの名前を更新する
	UpdateUserName(user model.User, userId uint) (model.UserResponse, error)
	// ユーザーを組織に加入させる
	AssignUserToOrganization(user model.User, userId uint) (model.UserAssignResponse, error)
	// ユーザーをチームに加入させる
	AssignUserToTeam(teamMember model.TeamMember, userId uint) (model.TeamMemberReponse, error)
	// ユーザーをチームから外す
	UnassignFromTeam(teamMember model.TeamMember, userId uint, teamId uint) (model.TeamMemberReponse, error)
	// 組織内のユーザー一覧情報を取得する
	GetOrganizationUsers(organizationId uint) ([]model.UserResponse, error)
}

type userUseCase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
	tmr repository.ITeamMemberRepository
}

func NewUserUseCase(ur repository.IUserRepository, uv validator.IUserValidator, tmr repository.ITeamMemberRepository) IUserUseCase {
	return &userUseCase{ur, uv, tmr}
}

func (uu *userUseCase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidator(user); err != nil {
		return model.UserResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash), Name: user.Name}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID: newUser.ID,
		Email: newUser.Email,
		Name: newUser.Name,
	}

	return resUser, err
}


func (uu *userUseCase) Login(user model.User) (string, error) {
	if err := uu.uv.UserValidator(user); err != nil {
		return "", err
	}
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func (uu *userUseCase) GetLoggedInUserDetails(user model.User, userId uint) (model.UserResponse, error) {
	if err := uu.ur.GetLoggedInUserDetails(&user, userId); err != nil {
		return model.UserResponse{}, nil
	}

	resUser := model.UserResponse {
		ID: user.ID,
		Email: user.Email,
		Name: user.Name,
	}

	return resUser, nil
}


func (uu *userUseCase) UpdateUserName(user model.User, userId uint) (model.UserResponse, error) {
	if err := uu.ur.UpdateUserName(&user, userId); err != nil {
		return model.UserResponse{}, err
	}

	resUser := model.UserResponse {
		ID: user.ID,
		Email: user.Email,
		Name: user.Name,
	}

	return resUser, nil
}


func (uu *userUseCase) AssignUserToOrganization(user model.User, userId uint) (model.UserAssignResponse, error) {
	if err := uu.ur.AssignUserToOrganization(&user, userId); err != nil {
		return model.UserAssignResponse{}, err
	}

	resUserAssign := model.UserAssignResponse {
		OrganizationId: user.OrganizationId,
	}

	return resUserAssign, nil
}


func (uu *userUseCase) AssignUserToTeam(teamMember model.TeamMember, userId uint) (model.TeamMemberReponse, error) {
	teamMembers := make([]model.TeamMember, 0)
	uu.tmr.GetTeamMembersByTeamId(&teamMembers, userId)
	for _, v := range teamMembers {
		if v.UserID == userId {
			return model.TeamMemberReponse{}, fmt.Errorf("the user is already assigned to the team")
		}
	}

	teamMember.UserID = userId
	if err := uu.tmr.AssignToTeam(&teamMember); err != nil {
		return model.TeamMemberReponse{}, err
	}

	resTeamMember := model.TeamMemberReponse {
		TeamID: teamMember.TeamID,
		UserID: teamMember.UserID,
		DeleteFlg: teamMember.DeleteFlg,
	}

	return resTeamMember, nil
}


func (uu *userUseCase) UnassignFromTeam(teamMember model.TeamMember, userId uint, teamId uint) (model.TeamMemberReponse, error) {
	if err := uu.tmr.UnassignFromTeam(&teamMember, userId, teamId); err != nil {
		return model.TeamMemberReponse{}, err
	}

	resTeamMember := model.TeamMemberReponse {
		TeamID: teamMember.TeamID,
		UserID: teamMember.UserID,
		DeleteFlg: teamMember.DeleteFlg,
	}

	return resTeamMember, nil
}


func (uu *userUseCase) GetOrganizationUsers(organizationId uint) ([]model.UserResponse, error) {
	users := make([]model.User, 0)
	if err := uu.ur.GetOrganizationUsers(&users, organizationId); err != nil {
		return nil, err
	}

	resUsers := make([]model.UserResponse, len(users))

	for i, v := range users {
		resUsers[i] = model.UserResponse{
			ID: v.ID,
			Email: v.Email,
			Name: v.Name,
		}
	}

	return resUsers, nil
}