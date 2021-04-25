package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EntityStudent struct {
	ID        string `gorm:"primaryKey;"`
	Name      string `gorm:"type:varchar(255);not null"`
	Npm       int    `gorm:"type:bigint;unique;not null"`
	Fak       string `gorm:"type:varchar(255);not null"`
	Bid       string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *EntityStudent) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityStudent) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
