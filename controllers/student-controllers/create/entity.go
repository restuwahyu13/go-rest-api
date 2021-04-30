package createStudent

type InputCreateStudent struct {
	Name string `json:"name" validate:"required,lowercase"`
	Npm  int    `json:"npm" validate:"required,number"`
	Fak  string `json:"fak" validate:"required,lowercase"`
	Bid  string `json:"bid" validate:"required,lowercase"`
}
