package routes

import (
	"todo-api-go/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterPagesRoutes(router *gin.RouterGroup) {

	webHandler := handler.NewWebHandler()

	// router.LoadHTMLGlob("internal/web/templates/*.html")

	router.GET("/", webHandler.HomePage)
}
