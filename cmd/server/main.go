package main

import (
	"todo-api-go/internal/api/middleware"
	"todo-api-go/internal/api/routes"
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

	router.Use(cors.Default())

	router.Use(middleware.Logger())

	store := sessions.NewCookieStore([]byte(config.SESSION_SECRET))
	store.MaxAge(config.SESSION_MAX_AGE)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = false
	store.Options.SameSite = 0

	gothic.Store = store

	googleProvider := google.New(
		config.GOOGLE_CLIENT_ID,
		config.GOOGLE_CLIENT_SECRET,
		config.GOOGLE_CALLBACK_URL,
	)

	goth.UseProviders(googleProvider)

	routes.RegisterAuthRoutes(router)
	routes.RegisterWebRoutes(router)
	routes.RegisterTodoRoutes(router)

	router.Run(config.APP_PORT)
}
