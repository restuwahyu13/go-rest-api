package resetAuth

type InputReset struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Cpassword string `json:"cpassword"`
	Active    bool   `json:"active"`
}
