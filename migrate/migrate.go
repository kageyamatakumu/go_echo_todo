package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
	"time"

	"gorm.io/gorm"
)

func main() {
	dbConn := db.CreateDB()
	defer fmt.Println("Successfully Migrated!")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(
		&model.User{},
		&model.Task{},
		&model.Organization{},
		&model.Team{},
		&model.InCharge{},
		&model.TeamMember{},
	)
	seed(dbConn)
}

func seed(db *gorm.DB) {
	unaffiliated := model.Organization{
		Name: "無所属",
		Description: "ユーザー作成後、初期に与えられる場所",
		Founder: 1,
	}
	organizaion := model.Organization{
		Name: "管理会社",
		Description: "初期のデータとして作成したもの",
		Founder: 1,
	}
	user := model.User{
		Name: "管理会社User",
		Email: "admin@sample.com",
		Password: "$2a$10$E/gZvqfuDl0LedbZymuj1uqyNIspFTLysSwov6EBi3XpZKn0CGuHa",
		OrganizationId: 2,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}
	db.Create(&unaffiliated)
	db.Create(&organizaion)
	db.Create(&user)
}