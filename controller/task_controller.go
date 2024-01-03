package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	GetTasksByDeadline(c echo.Context) error
	// CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	UpdateTaskStatus(c echo.Context) error
	DeleteTask(c echo.Context) error
	NarrowDownStatus(c echo.Context) error
	FuzzySearch(c echo.Context) error
}

type taskController struct {
	tu usecase.ITaskUseCase
}

func NewTaskController(tu usecase.ITaskUseCase) ITaskController {
	return &taskController{tu}
}

func (tc *taskController) GetAllTasks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	taskRes, err := tc.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) GetTaskById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)
	taskRes, err := tc.tu.GetTaskById(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, taskRes)
}

	func (tc *taskController) GetTasksByDeadline(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["user_id"]
		deadline_from , err := time.Parse("2006-01-02", c.QueryParam("deadline_from"))
		if err != nil {
			return err
		}
		deadline_to, err := time.Parse("2006-01-02", c.QueryParam("deadline_to"))
		if err != nil {
			return err
		}

		taskRes, err := tc.tu.GetTasksByDeadline(uint(userId.(float64)), deadline_from, deadline_to)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, taskRes)
	}

// func (tc *taskController) CreateTask(c echo.Context) error {
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	userId := claims["user_id"]

// 	task := model.Task{}
// 	if err := c.Bind(&task); err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	task.UserId = uint(userId.(float64))
// 	taskRes, err := tc.tu.CreateTask(task)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusCreated, taskRes)
// }

func (tc *taskController) UpdateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	taskRes, err := tc.tu.UpdateTask(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) UpdateTaskStatus(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	taskRes, err := tc.tu.UpdateTaskStatus(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, taskRes)
}


func (tc *taskController) DeleteTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	err := tc.tu.DeleteTask(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (tc *taskController) NarrowDownStatus(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	taskStatus := c.QueryParam("taskStatus")

	taskRes, err := tc.tu.NarrowDownStatus(uint(userId.(float64)), taskStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) FuzzySearch(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	tasStatus := c.QueryParam("taskStatus")
	search := c.QueryParam("search")

	taskRes, err := tc.tu.FuzzySearch(uint(userId.(float64)), search, tasStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, taskRes)
}