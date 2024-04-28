package main

import (
	"github.com/Rashad-Muntar/go-api/config"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadInitializers()
	config.ConnectDB()
}

func main() {
	r := gin.New()
	r.Run(":8080")
}