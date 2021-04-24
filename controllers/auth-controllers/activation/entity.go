package activationAuth

type InputActivation struct {
	Email  string `json:"email" binding:"required"`
	Active bool   `json:"active" binding:"required"`
	Token  string `json:"token" binding:"required"`
}
