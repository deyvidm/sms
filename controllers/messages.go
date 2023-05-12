package controllers

import (
	"net/http"

	"github.com/deyvidm/sms-backend/types"
	"github.com/gin-gonic/gin"
)

func NewMessage(c *gin.Context) {
	var input types.NewMessage
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}

	user, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}

	if err := user.SaveMessage(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": types.StatusSuccess})
}
