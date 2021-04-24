package model

import (
	"time"
)

type EntityRegister struct {
	ID        uint   `gorm:"type:bigserial;primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(255);not null"`
	Npm       uint32 `gorm:"type:bigint(255);unique;not null"`
	Fak       string `gorm:"type:varchar(255);not null"`
	Bid       string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
