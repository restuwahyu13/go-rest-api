package resetAuth

type InputReset struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8"`
	Cpassword string `json:"cpassword" validate:"required,gte=8"`
	Active    bool   `json:"active"`
}
