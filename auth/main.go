package main

import (
	"auth/db"
	"auth/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	db.Connect()
	r := gin.Default()

	r.POST("/signup", handler.Signup)
	r.POST("/login", handler.Signin)

	r.Run(":8080")
}
