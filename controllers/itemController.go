package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Rashad-Muntar/go-api/services"
	"github.com/gin-gonic/gin"
)

func ItemCreate(c *gin.Context) {
	var body struct {
		Price    float32 `json:"price" binding:"required"`
		Cost     float32 `json:"cost" binding:"required"`
		ItemName string  `json:"item_name" binding:"required"`
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

func ItemUpdate(c *gin.Context) {
	var body struct {
		ItemName string  `json:"item_name"`
		Cost     float32 `json:"cost"`
		Price    float32 `json:"price"`
		ItemID   uint    `json:"item_id" binding:"required"`
	}
	fmt.Println(body)
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse request body",
		})
		return
	}

	transact, err := services.ItemUpdate(c, body.ItemID, body.Price, body.Cost, body.ItemName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transact)
}

func ItemDelete(c *gin.Context) {
	itemID := c.Param("id")
	if itemID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Transaction ID is required",
		})
		return
	}
	id, err := strconv.ParseUint(itemID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid transaction ID",
		})
		return
	}
	err = services.ItemDelete(c, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete transaction",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction deleted successfully",
	})
}
