package routes

import (
	"github.com/Rashad-Muntar/go-api/controllers"
	"github.com/gin-gonic/gin"
)

func ItemRoute(router *gin.Engine) {
	router.POST("/item", controllers.ItemCreate)
	router.PUT("/item", controllers.ItemUpdate)
}