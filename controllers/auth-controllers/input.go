package controllers

type InputRegister struct {
	Fullname string `json:fullname binding:required`
	Email    string `json:"email" binding:"required"`
	Password uint   `json:"password" binding:"required"`
}

type InputLogin struct {
	Email    string `json:"email" binding:"required"`
	Password uint   `json:"password" binding:"required"`
}

type InputActivation struct {
	Token    string `uri:"token" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"email" binding:"required"`
}

type InputForgotPassword struct {
	Email string `json:"email"`
}

type inputResendActivation struct {
	Email string `json:"email"`
}

type InputResetPassword struct {
	Token     string `uri:"token" binding:"required"`
	password  string `json:"password" binding:"required"`
	cpassword string `json:"cpassword" binding:"required"`
}
