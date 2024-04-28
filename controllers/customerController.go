package controllers

import (
	"fmt"
	"net/http"

	"github.com/Rashad-Muntar/go-api/services"
	"github.com/gin-gonic/gin"
)


func CreateCustomerController(c *gin.Context) {
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
	user, err := services.CreateCustomerService(c, body.Balance, body.CustomerName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}
