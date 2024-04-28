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

type TransactionView struct {
	gorm.Model
    TransactionID uint
    CustomerID    uint
    CustomerName  string
    ItemName      string
    Qty           int32
    Price         float32
    Amount        float32
}