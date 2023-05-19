package controllers

import (
	"net/http"

	"github.com/deyvidm/sms-backend/models"
	"github.com/deyvidm/sms-backend/types"
	"github.com/gin-gonic/gin"
)

func ReceiveSNS(c *gin.Context) {
	var input types.SNSEvent

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}

	e, err := models.SNSEventFromInput(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err})
		return
	}

	_, err = e.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": "failed to save user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": types.StatusSuccess, "data": e})
}
