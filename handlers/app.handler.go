package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppHandler struct{}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (h *AppHandler) HomePage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", nil)
}
