package controllers

type Service interface {
	LoginService(payload)
}

type service struct {
	service Repository
}
