package main

import (
	"net/http"
	"todo-api-go/middlewares"
	"todo-api-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middlewares.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Todo Api",
		})
	})

	routes.RegisterTodoRoutes(router)

	router.Run()
}
