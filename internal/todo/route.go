package todo

import (
	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(router *gin.Engine) {
	todoHandler := NewTodoHandler()
	todoRoutes := router.Group("/todos")
	{
		todoRoutes.POST("/", todoHandler.CreateTodo)
		todoRoutes.GET("/", todoHandler.GetAllTodos)
		todoRoutes.GET("/:id", todoHandler.GetTodo)
		todoRoutes.PATCH("/:id", todoHandler.UpdateTodo)
		todoRoutes.DELETE("/:id", todoHandler.DeleteTodo)
	}
}
