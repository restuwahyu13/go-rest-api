package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EntityUsers struct {
	ID        string `gorm:"primaryKey;"`
	Fullname  string `gorm:"type:varchar(255);unique;not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Active    bool   `gorm:"type:bool;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *EntityUsers) BeforeCreate(db *gorm.DB) {
	entity.ID = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
}

func (entity *EntityUsers) BeforeUpdate(db *gorm.DB) {
	entity.UpdatedAt = time.Now().Local()
}
