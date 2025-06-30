package requests

type FormData struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type CreateUserRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
