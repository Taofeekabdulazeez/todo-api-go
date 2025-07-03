package routes

import (
	"todo-api-go/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAppRoutes(router *gin.Engine) {

	appHandler := handlers.NewAppHandler()

	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", appHandler.HomePage)
}
