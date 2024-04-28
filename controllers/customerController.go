package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Rashad-Muntar/go-api/services"
	"github.com/gin-gonic/gin"
)


func Create(c *gin.Context) {
	var body struct {
		CustomerName     string `json:"customer_name" binding:"required"`
		Balance    float32 `json:"balance" binding:"required"`
	}
	fmt.Println(body)
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	user, err := services.Create(c, body.Balance, body.CustomerName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	var body struct {
		CustomerID   uint    `json:"customer_id" binding:"required"`
		CustomerName string  `json:"customer_name"`
		Balance      float32 `json:"balance"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse request body",
		})
		return
	}

	customer, err := services.Update(c, body.CustomerID, body.Balance, body.CustomerName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func Delete(c *gin.Context) {
	customerID := c.Param("id")
	if customerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Customer ID is required",
		})
		return
	}

	id, err := strconv.ParseUint(customerID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid customer ID",
		})
		return
	}

	err = services.Delete(c, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete customer",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Customer deleted successfully",
	})
}
