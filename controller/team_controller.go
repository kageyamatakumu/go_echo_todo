package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITeamController interface {
	// 所属チームを取得する
	GetAssignTeamByUserId(c echo.Context) error
	// チームを作成する
	CreateTeam(c echo.Context) error
	// 組織内のチーム一覧を取得する
	GetTeamsByOrganizationId(c echo.Context) error
	// チームを削除する
	DeleteTeam(c echo.Context) error
}

type teamController struct {
	tu usecase.ITeamUseCase
}

func NewTeamController(tu usecase.ITeamUseCase) ITeamController {
	return &teamController{tu}
}

func (tc *teamController) GetAssignTeamByUserId(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	teamsRes, err := tc.tu.GetAssignTeamByUserId(uint(userId.(float64)));
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, teamsRes)
}

func (tc *teamController) CreateTeam(c echo.Context) error {
	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// userId := claims["user_id"]
	id := c.Param("organizationId")

	team := model.Team{}
	organizationId, _ := strconv.Atoi(id)
	team.OrganizationId = uint(organizationId)
	if err := c.Bind(&team); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	teamRes, err := tc.tu.CreateTeam(team)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, teamRes)
}

func (tc *teamController) GetTeamsByOrganizationId(c echo.Context) error {
	id := c.Param("organizationId")
	organizationId, _ := strconv.Atoi(id)

	teamRes, err := tc.tu.GetTeamsByOrganizationId(uint(organizationId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, teamRes)
}

func (tc *teamController) DeleteTeam(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("teamId")
	teamId, _ := strconv.Atoi(id)

	if err := tc.tu.DeleteTeam(uint(teamId), uint(userId.(float64))); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}