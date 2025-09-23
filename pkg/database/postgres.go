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
	var connectionErr error

	if DB, connectionErr = gorm.Open(postgres.Open(config.DSN), &gorm.Config{}); connectionErr != nil {
		log.Fatalln("Error connecting to database")
	}

	if migrationErr := DB.AutoMigrate(&model.User{}, &model.Todo{}); migrationErr != nil {
		log.Fatalln("Error performing database migrations")
	}
}
