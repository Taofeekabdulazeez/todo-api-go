package routes

import (
	"todo-api-go/internal/api/handler"
	"todo-api-go/internal/api/middleware"
	"todo-api-go/internal/core/service"

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
		authRoutes.POST("/logout", authHandler.Logout)
		authRoutes.GET("/sendMail", func(ctx *gin.Context) {
			emailS := &service.MailService{}
			if err := emailS.SendVerificationEmail("taofeekabdulazeeztaiwo@gmail.com", "hello"); err != nil {
				ctx.AbortWithStatusJSON(500, gin.H{
					"success": false,
					"message": "Email sending failed",
					"error":   err.Error(),
				})
				return
			}
			ctx.JSON(200, gin.H{"success": true, "message": "Email sent successfully"})
		})
	}
}
