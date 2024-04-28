package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ItemName string `gorm:"not null"`
	Cost float32`gorm:"not null"`
	Price float32`gorm:"not null"`
	Sort int
}