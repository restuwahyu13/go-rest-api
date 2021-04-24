package updateStudent

import "time"

type InputUpdateStudent struct {
	ID        uint64 `binding:"required"`
	Name      string `json:"name" binding:"required"`
	Npm       int    `json:"npm" binding:"required"`
	Fak       string `json:"fak" binding:"required"`
	Bid       string `json:"bid" binding:"required"`
	UpdatedAt time.Time
}
