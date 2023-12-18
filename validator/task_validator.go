package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
	TaskStatusValidate(task model.Task) error
}

type taskValidator struct {}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 10).Error("limited max 10 char"),
		),
	)
}

func (tv *taskValidator) TaskStatusValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Status,
			validation.In(model.TaskStatusUnstarted, model.TaskStatusStarted, model.TaskStatusCompleted).Error("The status must be one of the following: TaskStatusUnstarted, TaskStatusStarted, or TaskStatusCompleted."),
		),
	)
}
