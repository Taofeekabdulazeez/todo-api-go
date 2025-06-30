package requests

type CreateUserRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
