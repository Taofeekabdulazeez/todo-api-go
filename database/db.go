package database

import (
	"todo-api-go/config"
	"todo-api-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	c := config.Get()
	host := c.DB_HOST
	user := c.DB_USER
	password := c.DB_PASSWORD
	dbname := c.DB_NAME
	port := c.DB_PORT

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Todo{})
	DB = db
}
