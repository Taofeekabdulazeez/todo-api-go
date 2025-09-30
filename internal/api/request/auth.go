package request

type CreateUserRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type SignInRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type UpdateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
