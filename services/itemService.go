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

func ItemUpdate(c *gin.Context, item_id uint, price, cost float32, item_name string) (*models.Item, error) {
	var item models.Item
	result := config.DB.First(&item, item_id)
	if result.Error != nil {
		return nil, result.Error
	}

	item.Cost = cost
	item.Price = price
	item.ItemName = item_name

	result = config.DB.Save(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func ItemDelete(c *gin.Context, itemID uint) error {
	var item models.Item
	result := config.DB.First(&item, itemID)
	if result.Error != nil {
		return result.Error
	}

	result = config.DB.Delete(&item, itemID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}


