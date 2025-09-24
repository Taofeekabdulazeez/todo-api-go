package middleware

import (
	"errors"
	"net/http"
	"todo-api-go/internal/core/model"
	"todo-api-go/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func RequireAuth(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := getUserFromSession(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized User",
				"error":   err.Error(),
			})
		}

		ctx.Next()
	}
}

func getUserFromSession(ctx *gin.Context) (model.User, error) {
	session, err := gothic.Store.Get(ctx.Request, config.SESSION_KEY)
	if err != nil {
		return model.User{}, err
	}

	user, ok := session.Values["user"].(model.User)
	if !ok {
		return model.User{}, errors.New("no valid user in session")
	}

	return user, nil
}
