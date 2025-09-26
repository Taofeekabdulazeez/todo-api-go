package handler

import (
	"errors"
	"net/http"
	"todo-api-go/internal/api/request"
	"todo-api-go/internal/core/model"
	"todo-api-go/internal/core/service"
	"todo-api-go/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
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

	var found bool
	user, found := h.userService.GetUserByEmail(data.Email)

	if user != nil || found {
		var errorResponse string
		var statusCode int

		if found {
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
	user, err := getUserFromSession(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "No Authenticated User",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User fetched successfully",
		"data":    user,
	})

}

func (h *AuthHandler) SignUpWithGoogle(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	query.Add("provider", "google")
	ctx.Request.URL.RawQuery = query.Encode()

	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

func (h *AuthHandler) HandleGoogleAuth(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	query.Add("provider", "google")
	ctx.Request.URL.RawQuery = query.Encode()

	var session *sessions.Session
	var authUser goth.User
	var user *model.User
	var err error

	if authUser, err = gothic.CompleteUserAuth(ctx.Writer, ctx.Request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error authenticating user",
			"error":   err.Error(),
		})
		return
	}

	var found bool
	user, found = h.userService.GetUserByEmail(authUser.Email)

	if !found {
		if user, err = h.userService.CreateUserByEmail(authUser.Email); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error creating user",
				"error":   err.Error(),
			})
			return
		}
	}

	if session, err = gothic.Store.New(ctx.Request, config.SESSION_KEY); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to initialize session",
			"error":   err.Error(),
		})
		return
	}
	session.Values["user"] = user

	if err = session.Save(ctx.Request, ctx.Writer); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to save cookie",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User authenticated successfully",
		"data":    user,
	})
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	gothic.Logout(ctx.Writer, ctx.Request)
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
