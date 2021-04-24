package forgotAuth

type InputForgot struct {
	Email string `json:"email"  binding:"required"`
}
