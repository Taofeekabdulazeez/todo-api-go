package routes

import (
	"todo-api-go/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	authHandler := handlers.NewAuthHandler()
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/sign-up", authHandler.SignUp)
		authRoutes.POST("/sign-in", authHandler.SignIn)
		authRoutes.GET("/:email", authHandler.GetUser)
	}
}
