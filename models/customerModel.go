package models

import (
	"gorm.io/gorm"
	// "github.com/ethereum/go-ethereum/common"
)

type Customer struct {
	gorm.Model
	CustomerName string `gorm:"not null"`
	Balance float32`gorm:"not null"`
}