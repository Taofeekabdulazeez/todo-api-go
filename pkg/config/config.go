package config

import (
	"log"
	"os"
	"time"
	"todo-api-go/pkg/utils"

	"github.com/joho/godotenv"
)

var (
	APP_PORT               string
	DB_USER                string
	DB_PASSWORD            string
	DB_HOST                string
	DB_PORT                string
	DB_NAME                string
	DSN                    string
	GOOGLE_CLIENT_ID       string
	GOOGLE_CLIENT_SECRET   string
	GOOGLE_CALLBACK_URL    string
	GOOGLE_PROVIDER_SCOPES []string = []string{}
	SESSION_SECRET         string
	SESSION_KEY            string
	SESSION_MAX_AGE        int = int(12 * time.Hour)
	CLIENT_URL             string
	MAIL_HOST              string
	MAIL_PORT              int
	MAIL_USERNAME          string
	MAIL_PASSWORD          string
	JWT_SECRET             []byte
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
	SESSION_KEY = os.Getenv("SESSION_KEY")
	CLIENT_URL = os.Getenv("CLIENT_URL")
	MAIL_HOST = os.Getenv("MAIL_HOST")
	MAIL_PORT = utils.ParseInt(os.Getenv("MAIL_PORT"))
	MAIL_USERNAME = os.Getenv("MAIL_USERNAME")
	MAIL_PASSWORD = os.Getenv("MAIL_PASSWORD")
	JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))
}

func IsProd() bool {
	if os.Getenv("ENV") == "production" {
		return true
	} else {
		return false
	}

}
