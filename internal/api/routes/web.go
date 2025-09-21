package routes

import (
	"todo-api-go/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterWebRoutes(router *gin.Engine) {

	webHandler := handler.NewWebHandler()

	router.LoadHTMLGlob("internal/web/templates/*.html")

	router.GET("/", webHandler.HomePage)
}
