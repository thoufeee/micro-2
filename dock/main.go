package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"res": "HeLLo Docker"})
}

func main() {

	r := gin.Default()
	r.GET("/", Run)
	r.Run(":8080")
}
