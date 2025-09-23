package main

import (
	"time"
	"todo-api-go/internal/api/middleware"
	"todo-api-go/internal/api/routes"
	"todo-api-go/pkg/config"
	"todo-api-go/pkg/database"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	store := cookie.NewStore([]byte(config.SESSION_SECRET))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   24 * int(time.Hour),
		HttpOnly: true,
		Secure:   false,
	})
	gothic.Store = store
	router.Use(sessions.Sessions(gothic.SessionName, store))

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
