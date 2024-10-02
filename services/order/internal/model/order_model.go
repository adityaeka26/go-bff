package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"not null"`
	Item      string `gorm:"not null"`
	Quantity  int32  `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
