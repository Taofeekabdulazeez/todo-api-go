package routes

import (
	"todo-api-go/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/sign-up", handlers.SignUp)
		authRoutes.POST("/sign-in", handlers.SignIn)
	}
}
