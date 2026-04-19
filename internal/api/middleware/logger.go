package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		log.Printf("Request: %s %s\n", ctx.Request.Method, ctx.Request.URL.Path)
		ctx.Next()
	}
}
