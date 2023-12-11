package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
)

func main() {
	db := db.CreateDB()
	userRepository := repository.NewUserRepostory(db)
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUseCase(userRepository)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	userController := controller.NewUserContoller(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}