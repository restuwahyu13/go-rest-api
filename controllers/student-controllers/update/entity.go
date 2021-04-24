package updateStudent

import "time"

type InputUpdateStudent struct {
	Name     string `json:"name"`
	Npm      uint32 `json:"npm"`
	Fak      string `json:"fak"`
	Bid      string `json:"bid"`
	UpdateAt time.Time
}
