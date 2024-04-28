package routes

import (
	"github.com/Rashad-Muntar/go-api/controllers"
	"github.com/gin-gonic/gin"
)

func TransactionRoute(router *gin.Engine) {
	router.POST("/transaction", controllers.TransactionCreate)
	router.PUT("/transaction", controllers.TransactionUpdate)
	router.DELETE("/transaction/:id", controllers.TransactionDelete)
}