package activation

type InputActivation struct {
	Email  string `json:"email"`
	Active bool   `json:"active"`
	Token  string `json:"token"`
}
