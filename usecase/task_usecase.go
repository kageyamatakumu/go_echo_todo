package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
	"time"
)

type ITaskUseCase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (model.TaskResponse, error)
	GetTasksByDeadline(userId uint, fromDate time.Time, toDate time.Time) ([]model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	UpdateTaskStatus(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
	NarrowDownStatus(userId uint, taskStatus string) ([]model.TaskResponse, error)
	FuzzySearch(userId uint, keyword string, taskStatus ...string)([]model.TaskResponse, error)
}

type taskUseCase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUseCase {
	return &taskUseCase{tr, tv}
}

func (tu *taskUseCase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t := model.TaskResponse{
			ID: v.ID,
			Title: v.Title,
			DeadLine: v.DeadLine,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}

func (tu *taskUseCase) GetTaskById(userId uint, taskId uint) (model.TaskResponse, error) {
	task := model.Task{}
	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID: task.ID,
		Title: task.Title,
		DeadLine: task.DeadLine,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (tu *taskUseCase) GetTasksByDeadline(userId uint, fromDate time.Time, toDate time.Time) ([]model.TaskResponse, error) {
	tasks := make([]model.Task, 0)
	if err := tu.tr.GetTasksByDeadline(&tasks, userId, fromDate, toDate); err != nil {
		return nil, err
	}
	resTasks := make([]model.TaskResponse, len(tasks))
	for i, v := range tasks {
		resTasks[i] = model.TaskResponse {
			ID: v.ID,
			Title: v.Title,
			Status: v.Status,
			Memo: v.Memo,
			DeadLine: v.DeadLine,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return resTasks, nil
}

func (tu *taskUseCase) CreateTask(task model.Task) (model.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, nil
	}
	resTask := model.TaskResponse{
		ID: task.ID,
		Title: task.Title,
		Status: task.Status,
		Memo: task.Memo,
		DeadLine: task.DeadLine,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (tu *taskUseCase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID: task.ID,
		Title: task.Title,
		Status: task.Status,
		Memo: task.Memo,
		DeadLine: task.DeadLine,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (tu *taskUseCase) UpdateTaskStatus(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	if err := tu.tv.TaskStatusValidate(task); err != nil {
		return model.TaskResponse{}, err
	}
	if err := tu.tr.UpdateTaskStatus(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID: task.ID,
		Title: task.Title,
		Status: task.Status,
		Memo: task.Memo,
		DeadLine: task.DeadLine,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (tu *taskUseCase) DeleteTask(userId uint, taskId uint) error {
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}

	return nil
}

func (tu *taskUseCase) NarrowDownStatus(userId uint, taskStatus string) ([]model.TaskResponse, error) {
	var status int

	switch taskStatus {
		case "Unstarted":
			status = int(model.TaskStatusUnstarted)
		case "Started":
			status = int(model.TaskStatusStarted)
		case "Completed":
			status = int(model.TaskStatusCompleted)
	}

	tasks := []model.Task{}
	if err := tu.tr.NarrowDownStatus(&tasks, userId, status); err != nil {
		return nil , err
	}
	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t := model.TaskResponse{
			ID: v.ID,
			Title: v.Title,
			Status: v.Status,
			Memo: v.Memo,
			DeadLine: v.DeadLine,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}

	return resTasks, nil
}

func (tu *taskUseCase) FuzzySearch(userId uint, keyword string, taskStatus ...string)([]model.TaskResponse, error) {
	var status int

	switch taskStatus[0] {
		case "Unstarted":
			status = int(model.TaskStatusUnstarted)
		case "Started":
			status = int(model.TaskStatusStarted)
		case "Completed":
			status = int(model.TaskStatusCompleted)
		default:
			status = 99
	}

	tasks := make([]model.Task, 0)
	if status != 99 {
		if err := tu.tr.FuzzySearchStatus(&tasks, userId, keyword, status); err != nil {
			return nil, err
		}
	} else {
		if err := tu.tr.FuzzySearch(&tasks, userId, keyword); err != nil {
			return nil, err
		}
	}
	resTasks := make([]model.TaskResponse, len(tasks))
	for i, v := range tasks {
		resTasks[i] = model.TaskResponse {
			ID: v.ID,
			Title: v.Title,
			Status: v.Status,
			Memo: v.Memo,
			DeadLine: v.DeadLine,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return resTasks, nil
}