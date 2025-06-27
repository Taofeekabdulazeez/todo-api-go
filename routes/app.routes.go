package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAppRoutes(router *gin.Engine) {

	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})

	router.GET("/sign-up", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "sign-up.html", nil)
	})
}
