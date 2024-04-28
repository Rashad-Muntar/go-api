package main

import (
	"github.com/Rashad-Muntar/go-api/config"
	"github.com/Rashad-Muntar/go-api/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadInitializers()
	config.ConnectDB()
}

func main() {
	r := gin.New()
	routes.CustomerRoute(r)
	routes.TransactionRoute(r)
	r.Run(":8080")
}