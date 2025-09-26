package routes

import (
	"todo-api-go/internal/api/handler"
	"todo-api-go/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(router *gin.RouterGroup) {
	todoHandler := handler.NewTodoHandler()
	todoRoutes := router.Group("/todos", middleware.RequireAuth())
	{
		todoRoutes.POST("/", todoHandler.CreateTodo)
		todoRoutes.GET("/", todoHandler.GetAllTodos)
		todoRoutes.GET("/:id", todoHandler.GetTodo)
		todoRoutes.PATCH("/:id", todoHandler.UpdateTodo)
		todoRoutes.DELETE("/:id", todoHandler.DeleteTodo)
	}
}
