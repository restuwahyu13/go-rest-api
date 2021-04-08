package auth

import (
	"time"
)

type EntityUsers struct {
	ID        uint
	Fullname  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"unique;not null"`
	Active    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
