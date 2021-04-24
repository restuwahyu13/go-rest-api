package resetAuth

type InputReset struct {
	Email     string `json:"email"  binding:"required"`
	Password  string `json:"password"  binding:"required"`
	Cpassword string `json:"cpassword"  binding:"required"`
	Active    bool   `json:"active"  binding:"required"`
}
