package database

import (
	"todo-api-go/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (err error) {
	c := config.GetAll()
	DSN := c.DSN

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	return err
}
