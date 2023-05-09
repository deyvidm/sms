package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/deyvidm/sms-backend/auth"
	"github.com/deyvidm/sms-backend/models"
	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Username string `json:"username" binding:"required,alphanum,min=3,max=255"` // 3 is a holy number
	Password string `json:"password" binding:"required,alphanum,min=6,max=255"` // min 6 for brcypt hash
}

func Register(c *gin.Context) {
	var input LoginData

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
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("welcome %s!", input.Username)})
}

func Login(c *gin.Context) {
	var input LoginData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{
		Username: input.Username,
		Password: input.Password,
	}
	token, err := models.LoginUser(u.Username, u.Password)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect login details"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func CurrentUser(c *gin.Context) {
	userID, err := auth.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(userID)
	u.Password = "no :)"

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
