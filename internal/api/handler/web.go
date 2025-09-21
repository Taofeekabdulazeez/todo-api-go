package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebHandler struct{}

func NewWebHandler() WebHandler {
	return  WebHandler{}
}

func (h  WebHandler) HomePage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", nil)
}
