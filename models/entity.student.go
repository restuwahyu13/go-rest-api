package model

import (
	"time"
)

type EntityStudent struct {
	ID        uint64 `gorm:"type:bigserial;primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(255);not null"`
	Npm       int    `gorm:"type:bigint;unique;not null"`
	Fak       string `gorm:"type:varchar(255);not null"`
	Bid       string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
