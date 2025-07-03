package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	authHandler := NewAuthHandler()
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/sign-up", authHandler.SignUp)
		authRoutes.POST("/sign-in", authHandler.SignIn)
		authRoutes.GET("/:email", authHandler.GetUser)
	}
}
