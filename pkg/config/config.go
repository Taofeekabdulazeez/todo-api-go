package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	APP_PORT             string
	DB_USER              string
	DB_PASSWORD          string
	DB_HOST              string
	DB_PORT              string
	DB_NAME              string
	DSN                  string
	GOOGLE_CLIENT_ID     string
	GOOGLE_CLIENT_SECRET string
	GOOGLE_CALLBACK_URL  string
	SESSION_SECRET       string
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("An error loading .env file")
	}

	APP_PORT = ":" + os.Getenv("APP_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DSN = os.Getenv("DSN")
	GOOGLE_CLIENT_ID = os.Getenv("GOOGLE_CLIENT_ID")
	GOOGLE_CLIENT_SECRET = os.Getenv("GOOGLE_CLIENT_SECRET")
	GOOGLE_CALLBACK_URL = os.Getenv("GOOGLE_CALLBACK_URL")
	SESSION_SECRET = os.Getenv("SESSION_SECRET")
}

// type Config struct {
// 	APP_PORT             string
// 	DB_USER              string
// 	DB_PASSWORD          string
// 	DB_HOST              string
// 	DB_PORT              string
// 	DB_NAME              string
// 	DSN                  string
// 	GOOGLE_CLIENT_ID     string
// 	GOOGLE_CLIENT_SECRET string
// 	GOOGLE_CALLBACK_URL  string
// 	SESSION_SECRET       string
// }

// func GetAll() *Config {
// 	return &Config{
// 		APP_PORT:             ":" + os.Getenv("APP_PORT"),
// 		DB_USER:              os.Getenv("DB_USER"),
// 		DB_PASSWORD:          os.Getenv("DB_PASSWORD"),
// 		DB_HOST:              os.Getenv("DB_HOST"),
// 		DB_PORT:              os.Getenv("DB_PORT"),
// 		DB_NAME:              os.Getenv("DB_NAME"),
// 		DSN:                  os.Getenv("DSN"),
// 		GOOGLE_CLIENT_ID:     os.Getenv("GOOGLE_CLIENT_ID"),
// 		GOOGLE_CLIENT_SECRET: os.Getenv("GOOGLE_CLIENT_SECRET"),
// 		GOOGLE_CALLBACK_URL:  os.Getenv("GOOGLE_CALLBACK_URL"),
// 		SESSION_SECRET:       os.Getenv("SESSION_SECRET"),
// 	}
// }
