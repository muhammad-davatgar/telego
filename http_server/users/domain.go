package users

type SignUpRequest struct {
	Username string
	Email    string
	Password string
}

type ValidatedSignUp struct {
	Username string `json:"username" validate:"required,min=4,max=10"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
