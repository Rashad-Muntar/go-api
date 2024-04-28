package routes

import (
	"github.com/Rashad-Muntar/go-api/controllers"
	"github.com/gin-gonic/gin"
)

func CustomerRoute(router *gin.Engine) {
	router.POST("/customer", controllers.CreateCustomerController)
}