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

