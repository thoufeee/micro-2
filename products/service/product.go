package service

import (
	"net/http"
	"products/db"
	"products/model"

	"github.com/gin-gonic/gin"
)

func NewProduct(c *gin.Context) {
	var data struct {
		Name  string `json:"name"`
		Price string `json:"price"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "fill blanks"})
		return
	}

	new := &model.Product{
		Name:  data.Name,
		Price: data.Price,
	}

	if err := db.DB.Create(&new).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to create new product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"res": "successfuly created new user"})
}

// product list
func AllProduct(c *gin.Context) {
	var data []model.Product

	if err := db.DB.Find(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed"})
		return
	}

	c.JSON(http.StatusOK, data)
}
