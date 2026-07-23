package config

import (
	"fmt"
	"log"

	"fullstack-project/back-end/model"

	"github.com/glebarez/sqlite"
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	db, err := gorm.Open(sqlite.Open("automation.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database :", err)
	}

	DB = db

	err = DB.AutoMigrate(
		&model.User{},
		&model.Scrap{},
		&model.History{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database :", err)
	}

	fmt.Println("Database connected successfully")
}