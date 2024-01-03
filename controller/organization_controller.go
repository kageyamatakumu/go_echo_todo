package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IOrganizationController interface {
	// 作成した組織を表示する
	GetCreatedOrganizationsByUserId(c echo.Context) error
	// 組織を作成する
	CreateOrganization(c echo.Context) error
	// 組織の一覧を取得する
	ListOrganizations(c echo.Context) error
}

type organizationController struct {
	ou usecase.IOrganizationUseCase
}

func NewOrganizationController(ou usecase.IOrganizationUseCase) IOrganizationController {
	return &organizationController{ou}
}

func (oc *organizationController) GetCreatedOrganizationsByUserId(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	ResOrganizations, err := oc.ou.GetCreatedOrganizationsByUserId(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ResOrganizations)
}

func (oc *organizationController) CreateOrganization(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	organization := model.Organization{}
	organization.Founder = uint(userId.(float64))
	if err := c.Bind(&organization); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	organizationRes, err := oc.ou.CreateOrganization(organization)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, organizationRes)
}

func (oc *organizationController) ListOrganizations(c echo.Context) error {
	organizationRes, err := oc.ou.ListOrganizations();
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, organizationRes)
}