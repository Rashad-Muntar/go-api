package services

import (
	"github.com/Rashad-Muntar/go-api/config"
	"github.com/Rashad-Muntar/go-api/models"
	"github.com/gin-gonic/gin"
)

func TransactionCreate(c *gin.Context, customer_id uint, qty int32, amount, price float32, itemID uint) (*models.Transaction, error) {
	transaction := models.Transaction{CustomerID: customer_id, Qty: qty, Price: price, Amount: amount, ItemID: itemID}
	newTransact := config.DB.Create(&transaction)

	if newTransact.Error != nil {
		return nil, newTransact.Error
	}
	return &transaction, nil
}


func TransactionUpdate(c *gin.Context, transaction_id uint, qty int32, amount, price float32, itemID uint) (*models.Transaction, error) {
	var transact models.Transaction
	result := config.DB.First(&transact, transaction_id)
	if result.Error != nil {
		return nil, result.Error
	}

	transact.Qty = qty
	transact.Price = price
	transact.Amount = amount
	transact.ItemID = itemID

	result = config.DB.Save(&transact)
	if result.Error != nil {
		return nil, result.Error
	}

	return &transact, nil
}

func TransactionDelete(c *gin.Context, transactionID uint) error {
	var customer models.Transaction
	result := config.DB.First(&customer, transactionID)
	if result.Error != nil {
		return result.Error
	}

	result = config.DB.Delete(&customer, transactionID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}




