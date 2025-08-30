package main

import (
	"products/db"
	"products/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.Default()

	r.POST("/newproduct", service.NewProduct)
	r.GET("/allproduct", service.AllProduct)

	r.Run(":8081")
}
