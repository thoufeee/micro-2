package main

import (
	"orders/db"
	"orders/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.Default()

	r.POST("/createorder", service.Orders)
	r.GET("/allorder", service.Allorder)

	r.Run(":8082")
}
