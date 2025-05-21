package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func TodoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path + "endpoint invoked!")
		c.Next()
	}
}
