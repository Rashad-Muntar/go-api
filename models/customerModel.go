package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	CustomerName string `gorm:"not null"`
	Balance float32`gorm:"not null"`
	Transactions []Transaction
	
}