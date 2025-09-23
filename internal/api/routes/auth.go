package routes

import (
	"todo-api-go/internal/api/handler"
	"todo-api-go/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	authHandler := handler.NewAuthHandler()
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/sign-up", authHandler.SignUp)
		authRoutes.POST("/sign-in", authHandler.SignIn)
		authRoutes.GET("/me", authHandler.GetUserWithSession)
		authRoutes.GET("/user", middleware.RequireAuth(), authHandler.GetAuthUser)
		authRoutes.GET("/google/signup", authHandler.SignUpWithGoogle)
		authRoutes.GET("/google/callback", authHandler.HandleGoogleCallback)
	}
}
