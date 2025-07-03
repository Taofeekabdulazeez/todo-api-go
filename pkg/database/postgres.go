package database

import (
	"todo-api-go/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (err error) {
	c := config.GetAll()
	host := c.DB_HOST
	user := c.DB_USER
	password := c.DB_PASSWORD
	dbname := c.DB_NAME
	port := c.DB_PORT

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return err
}
