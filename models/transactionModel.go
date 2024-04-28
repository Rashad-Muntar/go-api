package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Qty int32 `gorm:"not null"`
	Price float32`gorm:"not null"`
	Amount float32`gorm:"not null"`
	CustomerID uint
	ItemID     uint
    Item       Item   `gorm:"foreignkey:ItemID"`
}