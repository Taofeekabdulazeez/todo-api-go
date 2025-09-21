package database

import (
	"log"
	"todo-api-go/internal/core/model"
	"todo-api-go/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	c := config.GetAll()
	DSN := c.DSN

	DB, connectionErr := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if connectionErr != nil {
		log.Fatalln("Error connecting to database")
	}

	migrationErr := DB.AutoMigrate(&model.User{}, &model.Todo{})

	if migrationErr != nil {
		log.Fatalln("Error performing database migrations")
	}
}
