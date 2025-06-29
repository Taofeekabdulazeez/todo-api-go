package handlers

import (
	"net/http"
	"strings"
	"todo-api-go/database"
	"todo-api-go/models"
	"todo-api-go/requests"
	"todo-api-go/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) SignUp(ctx *gin.Context) {

	var data requests.FormData

	err := ctx.ShouldBind(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Data",
			"data":  nil,
		})
		return
	}

	data.Email = strings.TrimSpace(data.Email)
	data.Password = strings.TrimSpace(data.Password)

	user := models.User{Email: data.Email, Password: data.Password}

	hash, err := utils.HashPassword(user.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"data":  nil,
		})
		return
	}

	user.Password = hash

	result := database.DB.Create(&user)

	if result.Error != nil {
		ctx.HTML(http.StatusBadRequest, "sign-up.html", nil)
		return
	}

	ctx.Redirect(http.StatusFound, "/")
}

func (h *AuthHandler) SignIn(ctx *gin.Context) {

}
