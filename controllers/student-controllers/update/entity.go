package createStudent

type InputCreateStudent struct {
	Name string `json:"name" binding:"required"`
	Npm  int    `json:"npm" binding:"required"`
	Fak  string `json:"fak" binding:"required"`
	Bid  string `json:"bid" binding:"required"`
}
