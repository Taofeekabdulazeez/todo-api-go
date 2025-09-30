package routes

import (
	"todo-api-go/internal/api/handler"
	"todo-api-go/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	authHandler := handler.NewAuthHandler()
	authRoutes := router.Group("/auth")
	{
		authRoutes.GET("/google/signup", authHandler.SignUpWithGoogle)
		authRoutes.GET("/google/callback", authHandler.HandleGoogleAuth)
		authRoutes.GET("/email/signup", authHandler.SignUpWithEmail)
		authRoutes.GET("/email/callback", authHandler.HandleEmailAuth)
		authRoutes.GET("/user", authHandler.GetAuthUser)
		authRoutes.PATCH("/user", middleware.RequireAuth(), authHandler.UpdateUser)
		authRoutes.GET("/logout", authHandler.Logout)
	}
}
