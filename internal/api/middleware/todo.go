package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func TodoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}
