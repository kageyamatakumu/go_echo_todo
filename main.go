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
	userUsecase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserContoller(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}