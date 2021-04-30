package forgotAuth

type InputForgot struct {
	Email string `json:"email" validate:"required,email"`
}
