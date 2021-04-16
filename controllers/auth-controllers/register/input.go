package register

type InputRegister struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type InputLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type InputActivation struct {
	Token    string `uri:"token" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type InputForgot struct {
	Email string `json:"email"`
}

type InputResend struct {
	Email string `json:"email"`
}

type InputResetPassword struct {
	Token     string `uri:"token" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Cpassword string `json:"cpassword" binding:"required"`
}
