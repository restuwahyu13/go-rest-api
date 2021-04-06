package controllers

type InputRegister struct {
	Name string `json:"name" binding:"required"`
	Npm  uint   `json:"npm" binding:"required"`
	Bid  string `json:"bid" binding:"required"`
	Fak  string `json:"fak" binding:"required"`
}
