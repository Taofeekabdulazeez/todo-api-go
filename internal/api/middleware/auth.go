package middleware

import (
	"net/http"
	"todo-api-go/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, err := gothic.Store.Get(ctx.Request, config.SESSION_KEY)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
				"err":     err.Error(),
			})
		}

		if user := session.Values["user"]; user == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized User",
			})
			return
		}
		ctx.Next()
	}
}
