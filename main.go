package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

func main() {
	db := db.CreateDB()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepostory(db)
	taskRepository := repository.NewTaskRepository(db)
	organizationRepository := repository.NewOrganizationRepository(db)
	teamRepository := repository.NewTeamRepository(db)
	teamMemberRepository := repository.NewTeamMemberRepository(db)
	userUsecase := usecase.NewUserUseCase(userRepository, userValidator, teamMemberRepository)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	organizationUsecase := usecase.NewOrganizationUseCase(organizationRepository)
	teamUsecase := usecase.NewTeamUseCase(teamRepository)
	userController := controller.NewUserContoller(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	organizationController := controller.NewOrganizationController(organizationUsecase)
	teamController := controller.NewTeamController(teamUsecase)
	e := router.NewRouter(userController, taskController, organizationController, teamController)
	e.Logger.Fatal(e.Start(":8080"))
}