package models

import (
	"gorm.io/gorm"
	// "github.com/ethereum/go-ethereum/common"
)

type Transaction struct {
	gorm.Model
	Qty int32 `gorm:"not null"`
	Price float32`gorm:"not null"`
	Amount float32`gorm:"not null"`
	CustomerID uint
}