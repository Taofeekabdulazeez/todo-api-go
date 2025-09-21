package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func AuthUser() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add("provider", "google")
		ctx.Request.URL.RawQuery = query.Encode()

		user, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
		println("User ===> ", user.Email)
		println("Error ===> ", err)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
