package requests

type FormData struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
