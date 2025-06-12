package database

import (
	"os"
	"todo-api-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {

	dbname := os.Getenv("dbname")
	host := os.Getenv("host")
	user := os.Getenv("user")
	password := os.Getenv("password")
	port := os.Getenv("port")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Todo{})
	DB = db

	return db
}
