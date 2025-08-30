package service

import (
	"net/http"
	"orders/db"
	"orders/model"

	"github.com/gin-gonic/gin"
)

// create order
func Orders(c *gin.Context) {
	var data struct {
		User_email   string `json:"email"`
		Product_Name string `json:"product"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid"})
		return
	}

	newoder := &model.Order{
		User_email:   data.User_email,
		Product_Name: data.Product_Name,
	}

	if err := db.DB.Create(&newoder).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"res": "successfuly orderd"})

}

// all order
func Allorder(c *gin.Context) {
	var data []model.Order

	db.DB.Find(&data)

	c.JSON(http.StatusOK, data)
}
