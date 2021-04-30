package activationAuth

type InputActivation struct {
	Email  string `json:"email" validate:"required,email"`
	Active bool   `json:"active"`
	Token  string `json:"token" validate:"required"`
}
