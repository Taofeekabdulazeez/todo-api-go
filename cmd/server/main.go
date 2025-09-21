package main

import (
	"todo-api-go/internal/api/middleware"
	"todo-api-go/internal/api/routes"
	"todo-api-go/pkg/config"
	"todo-api-go/pkg/database"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"

	"github.com/gin-gonic/gin"
)

// @title todo-api-go
// @version 1.0

func main() {
	config.Init()

	database.Connect()

	router := gin.Default()

	router.Use(middleware.Logger())

	googleProvider := google.New(
		config.GetAll().GOOGLE_CLIENT_ID,
		config.GetAll().GOOGLE_CLIENT_SECRET,
		config.GetAll().GOOGLE_CALLBACK_URL,
	)

	goth.UseProviders(googleProvider)

	routes.RegisterAuthRoutes(router)
	routes.RegisterWebRoutes(router)
	routes.RegisterTodoRoutes(router)

	router.Run(":8080")
}
