package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		fmt.Println("LoggerMiddleware invoked!")
		ctx.Next()
	}
}
