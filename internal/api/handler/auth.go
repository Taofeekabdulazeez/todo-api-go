package handler

import (
	"errors"
	"net/http"
	"time"
	"todo-api-go/internal/api/request"
	"todo-api-go/internal/core/model"
	"todo-api-go/internal/core/service"
	"todo-api-go/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

type AuthHandler struct {
	userService         *service.UserService
	verificationService *service.VerificationService
	mailService         *service.MailService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		userService:         &service.UserService{},
		verificationService: &service.VerificationService{},
		mailService:         &service.MailService{},
	}
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

	if err := storeSession(ctx, "user", user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to store user session",
			"error":   err.Error(),
		})
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, config.CLIENT_URL)

}

func (h *AuthHandler) SignUpWithEmail(ctx *gin.Context) {
	var email string
	if email = ctx.Query("email"); email == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Email is required",
			"error":   "missing email query parameter",
		})
		return
	}

	var found bool
	var user *model.User
	var err error

	user, found = h.userService.GetUserByEmail(email)
	if !found {
		if user, err = h.userService.CreateUserByEmail(email); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error creating user",
				"error":   err.Error(),
			})
			return
		}
	}

	token, err := h.verificationService.CreateToken(user.Email, time.Hour)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error verifying user",
			"error":   err.Error(),
		})
		return
	}

	if err := h.mailService.SendVerificationEmail(email, token); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error sending verification email",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Verification email sent",
	})
}

func (h *AuthHandler) HandleEmailAuth(ctx *gin.Context) {
	var token string
	if token = ctx.Query("token"); token == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Token is required",
			"error":   "missing token query parameter",
		})
		return
	}

	email, err := h.verificationService.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error verifying token",
			"error":   err.Error(),
		})
		return
	}

	user, found := h.userService.GetUserByEmail(email.(string))

	if !found {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "No user found for this token",
			"error":   "user not found",
		})
		return
	}

	if err := storeSession(ctx, "user", user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to store user session",
			"error":   err.Error(),
		})
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, config.CLIENT_URL)
}

func (h *AuthHandler) GetAuthUser(ctx *gin.Context) {
	sessionUser, err := getUserFromSession(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "No Authenticated User",
			"error":   err.Error(),
		})
		return
	}

	user, found := h.userService.GetUserByEmail(sessionUser.Email)
	if !found {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "No Authenticated User",
			"error":   "User retrieving user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User fetched successfully",
		"data":    user,
	})

}

func (h *AuthHandler) UpdateUser(ctx *gin.Context) {
	var body request.UpdateUserRequest
	user, _ := ctx.Get("user")

	if validationErr := ctx.ShouldBindBodyWithJSON(&body); validationErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid Request Data",
			"error":   validationErr.Error(),
		})
		return
	}

	updatedUser, err := h.userService.UpdateUser(user.(model.User), body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error updating user data",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User data updated successfully",
		"data":    updatedUser,
	})
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	if err := gothic.Logout(ctx.Writer, ctx.Request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Couldn't log out user",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User successfully logged out",
		"data":    nil,
	})
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

func storeSession(ctx *gin.Context, key string, value any) error {

	session, err := gothic.Store.New(ctx.Request, config.SESSION_KEY)
	if err != nil {
		return err
	}

	session.Values[key] = value

	if err = session.Save(ctx.Request, ctx.Writer); err != nil {
		return err
	}

	return nil
}
