package repository

import (
	"fmt"
	"go-rest-api/model"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	GetTasksByDeadline(task *[]model.Task, userId uint, fromDate time.Time, toDate time.Time) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	UpdateTaskStatus(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
	NarrowDownStatus(tasks *[]model.Task, userId uint, taskStatus int) error
	FuzzySearch(tasks *[]model.Task, userId uint, keyword string) error
	FuzzySearchStatus(tasks *[]model.Task, userId uint, keyword string, taskStatus int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}


func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	// SELECT * FROM tasks LEFT JOIN users ON tasks.user_id = users.id WHERE tasks.user_id = {userId} ORDER BY task.created_at desc;
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}

	return nil
}

func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	// SELECT * FROM tasks LEFT JOIN users ON tasks.user_id = users.id WHERE tasks.user_id = {userId}, id = {taskId} ORDER BY id LIMIT 1;
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(task, taskId).Error; err != nil {
		return err
	}

	return nil
}

func (tr *taskRepository) GetTasksByDeadline(tasks *[]model.Task, userId uint, fromDate time.Time, toDate time.Time) error {
	if err := tr.db.Joins("User").Where("user_id=? AND dead_line BETWEEN ? AND ?", userId, fromDate, toDate).Find(tasks).Error; err != nil {
		return err
	}

	return nil
}

func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}

	return nil
}

func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskId, userId).Updates(map[string]interface{}{"title": task.Title, "memo": task.Memo, "status": task.Status, "dead_line": task.DeadLine})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *taskRepository) UpdateTaskStatus(task *model.Task, userId uint, taskId uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskId, userId).Update("status", task.Status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *taskRepository) NarrowDownStatus(tasks *[]model.Task, userId uint, taskStatus int) error {
	if err := tr.db.Joins("User").Where("user_id=? AND status=?", userId, taskStatus).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) FuzzySearch(tasks *[]model.Task, userId uint, keyword string) error {
	if err := tr.db.Joins("User").Where("user_id=? AND (title LIKE ? OR memo LIKE ?)", userId, "%"+keyword+"%", "%"+keyword+"%").Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) FuzzySearchStatus(tasks *[]model.Task, userId uint, keyword string, taskStatus int) error {
	if err := tr.db.Joins("User").Where("user_id=? AND (title LIKE ? OR memo LIKE ?) AND status=? ", userId, "%"+keyword+"%", "%"+keyword+"%", taskStatus).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}