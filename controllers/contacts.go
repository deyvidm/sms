package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Contacts(c *gin.Context) {

}

type NewContactData struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

func NewContact(c *gin.Context) {
	var input NewEventData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "validated!", "data": input})
}
