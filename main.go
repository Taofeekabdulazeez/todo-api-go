package main

import (
	"net/http"
	"todo-api-go/controllers"
	"todo-api-go/middlewares"

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

	router.GET("/todos", controllers.GetAllTodos)
	router.POST("/todos", controllers.CreateTodo)

	todoRoutes := router.Group("/todos").Use(middlewares.TodoMiddleware())
	{
		todoRoutes.POST("/", controllers.CreateTodo)
		todoRoutes.GET("/", controllers.GetAllTodos)
		todoRoutes.GET("/:id", controllers.GetTodo)
		todoRoutes.PATCH("/:id", controllers.UpdateTodo)
		todoRoutes.DELETE("/:id", controllers.DeleteTodo)
	}

	router.Run()
}
