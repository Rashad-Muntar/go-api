package services

import (
	"github.com/Rashad-Muntar/go-api/config"
	"github.com/Rashad-Muntar/go-api/models"
	"github.com/gin-gonic/gin"
)

func ItemCreate(c *gin.Context, price, cost float32, item_name string) (*models.Item, error) {
	item := models.Item{ItemName: item_name, Cost: cost, Price: price}
	newCustomer := config.DB.Create(&item)

	if newCustomer.Error != nil {
		return nil, newCustomer.Error
	}
	return &item, nil
}
