package routes

import (
	"todo-api-go/controllers"
	"todo-api-go/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(router *gin.Engine) {
	todoRoutes := router.Group("/todos").Use(middlewares.TodoMiddleware())
	{
		todoRoutes.POST("/", controllers.CreateTodo)
		todoRoutes.GET("/", controllers.GetAllTodos)
		todoRoutes.GET("/:id", controllers.GetTodo)
		todoRoutes.PATCH("/:id", controllers.UpdateTodo)
		todoRoutes.DELETE("/:id", controllers.DeleteTodo)
	}
}
