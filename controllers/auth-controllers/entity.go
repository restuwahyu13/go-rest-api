package auth

import (
	"time"
)

type EntityUsers struct {
	ID        uint   `gorm:"type:bigserial;primaryKey;autoIncrement"`
	Fullname  string `gorm:"type:varchar(255);unique;not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Active    bool   `gorm:"type:bool;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
