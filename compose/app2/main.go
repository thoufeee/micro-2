package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run(c *gin.Context) {

	c.JSON(http.StatusOK, "HELLOOO DOCKERRR 222222")
}

func main() {
	r := gin.Default()

	r.GET("/app2", Run)

	r.Run(":8081")
}
