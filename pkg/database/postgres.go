package database

import (
	"log"
	"todo-api-go/internal/core/model"
	"todo-api-go/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	var connectionErr error

	if DB, connectionErr = gorm.Open(postgres.Open(config.DSN), &gorm.Config{PrepareStmt: false, Logger: logger.Default.LogMode(logger.Silent)}); connectionErr != nil {
		log.Fatalln("Error connecting to database", connectionErr)
	}

	if migrationErr := DB.AutoMigrate(&model.User{}, &model.Todo{}); migrationErr != nil {
		log.Fatalln("Error performing database migrations", migrationErr)
	}
}
