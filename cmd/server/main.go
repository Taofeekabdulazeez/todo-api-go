package main

import (
	"todo-api-go/internal/api/middleware"
	"todo-api-go/internal/api/routes"
	"todo-api-go/internal/core/model"
	"todo-api-go/pkg/config"
	"todo-api-go/pkg/database"

	"github.com/gin-gonic/gin"
)

// @title todo-api-go
// @version 1.0

func main() {
	config.Init()

	err := database.Connect()
	if err != nil {
		panic(err)
	}

	database.DB.AutoMigrate(&model.User{}, &model.Todo{})

	router := gin.Default()

	router.Use(middleware.Logger())

	routes.RegisterAuthRoutes(router)
	routes.RegisterWebRoutes(router)
	routes.RegisterTodoRoutes(router)

	router.Run(":8080")
}
