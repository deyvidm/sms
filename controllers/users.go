package controllers

import (
	"net/http"

	"github.com/deyvidm/sms-backend/models"
	"github.com/gin-gonic/gin"
)

type SignInData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input SignInData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{
		Username: input.Username,
		Password: input.Password,
	}
	_, err := u.SaveUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed to save user"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "validated!", "data": input})
}

func SignIn(c *gin.Context) {
	var input SignInData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect login details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
