package main

import (
	"net/http"
	"todo-api-go/configs"
	"todo-api-go/database"
	"todo-api-go/middlewares"
	"todo-api-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.Init()
	router := gin.Default()
	router.Use(middlewares.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Todo Api",
		})
	})

	database.Connect()

	routes.RegisterTodoRoutes(router)

	router.Run()
}
