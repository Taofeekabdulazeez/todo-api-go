package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetAll() *Config {

	return &Config{
		DB_USER:     os.Getenv("user"),
		DB_PASSWORD: os.Getenv("password"),
		DB_HOST:     os.Getenv("host"),
		DB_PORT:     os.Getenv("port"),
		DB_NAME:     os.Getenv("dbname"),
	}
}
