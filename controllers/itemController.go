package controllers

import (
	"net/http"

	"github.com/Rashad-Muntar/go-api/services"
	"github.com/gin-gonic/gin"
)

func ItemCreate(c *gin.Context) {
	var body struct {
		Price       float32 `json:"price" binding:"required"`
		Cost     	float32 `json:"cost" binding:"required"`
		ItemName    string   `json:"item_name" binding:"required"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	user, err := services.ItemCreate(c, body.Price, body.Cost, body.ItemName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}
