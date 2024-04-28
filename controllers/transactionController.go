package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Rashad-Muntar/go-api/services"
	"github.com/gin-gonic/gin"
)

func TransactionCreate(c *gin.Context) {
	var body struct {
		Price      float32 `json:"price" binding:"required"`
		Amount     float32 `json:"amount" binding:"required"`
		Qty        int32   `json:"qty" binding:"required"`
		CustomerID uint    `json:"customer_id" binding:"required"`
	}
	fmt.Println(body)
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	user, err := services.TransactionCreate(c, body.CustomerID, body.Qty, body.Amount, body.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}


func TransactionUpdate(c *gin.Context) {
	var body struct {
		TransactionID uint `json:"transaction_id" binding:"required"`
		Price  float32 `json:"price"`
		Amount float32 `json:"amount"`
		Qty    int32   `json:"qty"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse request body",
		})
		return
	}

	transact, err := services.TransactionUpdate(c, body.TransactionID, body.Qty, body.Amount, body.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transact)
}


func TransactionDelete(c *gin.Context) {
	TransactionID := c.Param("id")
	if TransactionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Transaction ID is required",
		})
		return
	}
	id, err := strconv.ParseUint(TransactionID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid transaction ID",
		})
		return
	}
	err = services.TransactionDelete(c, uint(id))
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