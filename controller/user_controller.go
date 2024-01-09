package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
	UpdateUserName(c echo.Context) error
	// ログインしているユーザーの情報を取得
	GetLoggedInUserDetails(c echo.Context) error
	// ユーザーを組織に加入させる
	AssignUserToOrganization(c echo.Context) error
	// ユーザーをチームに加入させる
	AssignUserToTeam(c echo.Context) error
	// ユーザーをチームから外す
	UnassignFromTeam(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUseCase
}

func NewUserContoller(uu usecase.IUserUseCase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwtToken"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "jwtToken"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}


func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}


func (uc *userController) UpdateUserName(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	userModel := model.User{}
	if err := c.Bind(&userModel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userRes, err := uc.uu.UpdateUserName(userModel, uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, userRes)
}

func (uc *userController) GetLoggedInUserDetails(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	userModel := model.User{}
	userRes, err := uc.uu.GetLoggedInUserDetails(userModel, uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, userRes)
}

func (uc *userController) AssignUserToOrganization(c echo.Context) error {
	userModel := model.User{}
	if err := c.Bind(&userModel); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userRes, err := uc.uu.AssignUserToOrganization(userModel, userModel.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, userRes)
}


func (uc *userController) AssignUserToTeam(c echo.Context) error {
	teamMember := model.TeamMember{}
	if err := c.Bind(&teamMember); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	teamMemberRes, err := uc.uu.AssignUserToTeam(teamMember, teamMember.UserID);
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, teamMemberRes)
}


func (uc *userController) UnassignFromTeam(c echo.Context) error {
	teamMember := model.TeamMember{}
	if err := c.Bind(&teamMember); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	teamMemberRes, err := uc.uu.UnassignFromTeam(teamMember, teamMember.UserID, teamMember.TeamID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, teamMemberRes)
}