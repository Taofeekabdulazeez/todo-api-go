package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *UserService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		userService: &UserService{},
	}
}

func (h *AuthHandler) SignUp(ctx *gin.Context) {
	var data CreateUserRequest

	err := ctx.ShouldBind(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Data",
			"data":  nil,
		})
		return
	}

	user, err := h.userService.GetUserByEmail(data.Email)

	if user != nil || err != nil {
		var errorResponse string
		var statusCode int

		if user != nil {
			errorResponse = "Account already Exist"
			statusCode = http.StatusBadRequest
		} else {
			errorResponse = "Internal Server error"
			statusCode = http.StatusInternalServerError
		}

		ctx.JSON(statusCode, gin.H{
			"error": errorResponse,
			"data":  nil,
		})
		return
	}

	user, err = h.userService.CreateUser(data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"data":  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User successfully signed up",
		"data":    user,
	})
}

func (h *AuthHandler) SignIn(ctx *gin.Context) {
	var data SignInRequest

	err := ctx.ShouldBind(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Data",
			"data":  nil,
		})
		return
	}

	success, user, err := h.userService.VerifyUser(data.Email, data.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"data":  nil,
		})
		return
	}

	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Sign In credentials",
			"data":  nil,
		})
		return
	}

	// ctx.HTML(http.StatusOK, "home.html", user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
	})
}

func (h *AuthHandler) GetUser(ctx *gin.Context) {
	email := ctx.Param("email")

	user, err := h.userService.GetUserByEmail(email)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
			"data":  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
	})
}
