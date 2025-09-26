package routes

import (
	"todo-api-go/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	authHandler := handler.NewAuthHandler()
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/sign-up", authHandler.SignUp)
		authRoutes.POST("/sign-in", authHandler.SignIn)
		authRoutes.GET("/user", authHandler.GetAuthUser)
		authRoutes.GET("/google/signup", authHandler.SignUpWithGoogle)
		authRoutes.GET("/google/callback", authHandler.HandleGoogleAuth)
	}
}
