package registerAuth

type InputRegister struct {
	Fullname string `json:"fullname" validate:"required,lowercase"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}
