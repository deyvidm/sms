package controllers

import (
	"net/http"

	"github.com/deyvidm/sms-backend/models"
	"github.com/gin-gonic/gin"
)

type NewContactData struct {
	FirstName string `json:"first_name" binding:"required,alpha,min=3,max=50"`
	LastName  string `json:"last_name" binding:"required,alpha,min=3,max=50"`
	Phone     string `json:"phone" binding:"required,e164"` // e164 is the standard +11234567890
}

func AllContacts(c *gin.Context) {
	user, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contacts, err := user.AllContacts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": contacts})
}

func NewContact(c *gin.Context) {
	var input NewContactData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact := models.Contact{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.Phone,
	}

	if _, err := user.SaveContact(contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": input})
}
