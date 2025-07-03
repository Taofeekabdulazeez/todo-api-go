package web

import (
	"github.com/gin-gonic/gin"
)

func RegisterAppRoutes(router *gin.Engine) {

	appHandler := NewAppHandler()

	router.LoadHTMLGlob("internal/web/templates/*.html")

	router.GET("/", appHandler.HomePage)
}
