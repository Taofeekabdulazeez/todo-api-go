package main

import (
	"todo-api-go/config"
	"todo-api-go/database"
	"todo-api-go/middlewares"
	"todo-api-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()

	router := gin.Default()

	router.Use(middlewares.Logger())

	routes.RegisterAppRoutes(router)
	routes.RegisterAuthRoutes(router)
	routes.RegisterTodoRoutes(router)

	database.Connect()

	router.Run()
}
