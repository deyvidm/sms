package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewEventData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetEvents(c *gin.Context) {

}

func NewEvent(c *gin.Context) {
	var input NewEventData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "validated!", "data": input})
}
