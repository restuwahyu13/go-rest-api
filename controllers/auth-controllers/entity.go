package controllers

import (
	"time"
)

type EntityRegister struct {
	ID        uint
	Name      string `gorm:"not null"`
	Npm       uint32 `gorm:"unique;not null"`
	Fak       string `gorm:"not null"`
	Bid       string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
