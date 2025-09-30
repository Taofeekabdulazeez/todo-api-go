package main

import (
	"encoding/gob"
	"net/http"
	"time"
	"todo-api-go/internal/api/routes"
	"todo-api-go/internal/core/model"
	"todo-api-go/pkg/config"
	"todo-api-go/pkg/database"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()

	database.Connect()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:9000",
			"https://mytodoit.vercel.app",
			"https://mytodoit.lovable.app",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	store := sessions.NewCookieStore([]byte(config.SESSION_SECRET))
	store.MaxAge(config.SESSION_MAX_AGE)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = config.IsProd()
	store.Options.SameSite = http.SameSiteNoneMode

	gothic.Store = store

	googleProvider := google.New(
		config.GOOGLE_CLIENT_ID,
		config.GOOGLE_CLIENT_SECRET,
		config.GOOGLE_CALLBACK_URL,
		config.GOOGLE_PROVIDER_SCOPES...,
	)

	goth.UseProviders(googleProvider)

	gob.Register(model.User{})

	api := router.Group("/api/v1")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	routes.RegisterAuthRoutes(api)
	routes.RegisterTodoRoutes(api)
	routes.RegisterPagesRoutes(api)

	router.Run(config.APP_PORT)
}
