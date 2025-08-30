package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run(c *gin.Context) {

	c.JSON(http.StatusOK, "HELLOOOOO DOCKER  1111")
}

func main() {
	r := gin.Default()

	r.GET("/app", Run)

	r.Run(":8080")
}
