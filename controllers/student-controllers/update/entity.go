package updateStudent

import (
	"time"
)

type InputUpdateStudent struct {
	ID        string `validate:"required,uuid"`
	Name      string `json:"name" validate:"required,lowercase"`
	Npm       int    `json:"npm" validate:"required,number"`
	Fak       string `json:"fak" validate:"required,lowercase"`
	Bid       string `json:"bid" validate:"required,lowercase"`
	UpdatedAt time.Time
}
