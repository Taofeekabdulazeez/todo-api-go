package main

import (
	"todo-api-go/internal/auth"
	"todo-api-go/internal/middleware"
	"todo-api-go/internal/todo"
	"todo-api-go/internal/web"
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

	database.DB.AutoMigrate(&auth.User{}, &todo.Todo{})

	router := gin.Default()

	router.Use(middleware.Logger())

	web.RegisterAppRoutes(router)
	auth.RegisterAuthRoutes(router)
	todo.RegisterTodoRoutes(router)

	router.Run(":8080")
}
