package routes

import (
	"todo-api-go/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(router *gin.Engine) {
	todoRoutes := router.Group("/todos")
	{
		todoRoutes.POST("/", handlers.CreateTodo)
		todoRoutes.GET("/", handlers.GetAllTodos)
		todoRoutes.GET("/:id", handlers.GetTodo)
		todoRoutes.PATCH("/:id", handlers.UpdateTodo)
		todoRoutes.DELETE("/:id", handlers.DeleteTodo)
	}
}
