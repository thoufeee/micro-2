package handler

import (
	"auth/db"
	"auth/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// signup
func Signup(c *gin.Context) {

	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "fill blanks"})
		return
	}

	new := &model.User{
		Email:    data.Email,
		Password: data.Password,
	}

	if err := db.DB.Create(&new).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to create new user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"res": "successfuly created new user"})
}

// signin
func Signin(c *gin.Context) {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "fill all blanks"})
		return
	}

	var user model.User

	if err := db.DB.Where("email = ?", &data.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "inavlid email or password"})
		return
	}

	if err := db.DB.Where("password = ?", &data.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"res": "successfily loged in"})
}
