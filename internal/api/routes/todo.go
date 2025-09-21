package routes

import (
	"todo-api-go/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(router *gin.Engine) {
	todoHandler := handler.NewTodoHandler()
	todoRoutes := router.Group("/todos")
	{
		todoRoutes.POST("/", todoHandler.CreateTodo)
		todoRoutes.GET("/", todoHandler.GetAllTodos)
		todoRoutes.GET("/:id", todoHandler.GetTodo)
		todoRoutes.PATCH("/:id", todoHandler.UpdateTodo)
		todoRoutes.DELETE("/:id", todoHandler.DeleteTodo)
	}
}
