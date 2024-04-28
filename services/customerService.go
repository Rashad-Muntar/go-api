package services

import (
	"github.com/Rashad-Muntar/go-api/config"
	"github.com/Rashad-Muntar/go-api/models"
	"github.com/gin-gonic/gin"
)

func CreateCustomerService(c *gin.Context, balance float32, customer_name string) (*models.Customer, error) {
	customer := models.Customer{Balance: balance, CustomerName: customer_name}
	newCustomer := config.DB.Create(&customer)

	if newCustomer.Error != nil {
		return nil, newCustomer.Error
	}
	return &customer, nil
}

func UpdateCustomerService(c *gin.Context, customerID uint, balance float32, customerName string) (*models.Customer, error) {
	var customer models.Customer
	result := config.DB.First(&customer, customerID)
	if result.Error != nil {
		return nil, result.Error
	}

	customer.Balance = balance
	customer.CustomerName = customerName

	result = config.DB.Save(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

