package controllers

import (
	"net/http"

	"github.com/deyvidm/sms-backend/models"
	"github.com/deyvidm/sms-backend/types"
	"github.com/gin-gonic/gin"
)

func UpdateInvite(c *gin.Context) {
	id := c.Param("id")
	var input types.UpdateInvite

	invite, err := models.GetInvite(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}

	err = invite.Save(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": types.StatusSuccess, "data": input})
}
