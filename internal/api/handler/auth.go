package handler

import (
	"errors"
	"net/http"
	"todo-api-go/internal/api/request"
	"todo-api-go/internal/core/service"
	"todo-api-go/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		userService: &service.UserService{},
	}
}

func (h *AuthHandler) SignUp(ctx *gin.Context) {
	var data request.CreateUserRequest

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
	var data request.SignInRequest

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

func (h *AuthHandler) GetAuthUser(ctx *gin.Context) {
	user, exist := ctx.Get("user")
	if !exist {
		ctx.AbortWithError(http.StatusNotFound, errors.New("user not found"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    user,
	})

}

func (h *AuthHandler) SignUpWithGoogle(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	query.Add("provider", "google")
	ctx.Request.URL.RawQuery = query.Encode()

	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

func (h *AuthHandler) HandleGoogleCallback(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	query.Add("provider", "google")
	ctx.Request.URL.RawQuery = query.Encode()

	user, _ := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)

	session, err := gothic.Store.New(ctx.Request, config.SESSION_KEY)

	session.Values["user"] = user

	if saveError := session.Save(ctx.Request, ctx.Writer); saveError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "failed to save cookie",
			"error":   saveError,
		})
		return
	}

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    user,
	})
}

func (h *AuthHandler) GetUserWithSession(ctx *gin.Context) {
	session, err := gothic.Store.Get(ctx.Request, config.SESSION_KEY)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}

	user := session.Values["user"]
	ctx.JSON(200, gin.H{"data": user})
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	gothic.Logout(ctx.Writer, ctx.Request)
}
