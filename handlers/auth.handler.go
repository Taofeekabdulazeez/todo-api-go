package handlers

import (
	"net/http"
	"todo-api-go/requests"
	"todo-api-go/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *services.UserService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		userService: &services.UserService{},
	}
}

func (h *AuthHandler) SignUp(ctx *gin.Context) {

	var data requests.CreateUserRequest

	err := ctx.ShouldBind(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Data",
			"data":  nil,
		})
		return
	}

	user, err := h.userService.CreateUser(data)

	if err != nil {
		ctx.HTML(http.StatusBadRequest, "sign-up.html", user)
		return
	}

	ctx.Redirect(http.StatusFound, "/")
}

func (h *AuthHandler) SignIn(ctx *gin.Context) {

}
