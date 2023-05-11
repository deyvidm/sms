package controllers

import (
	"net/http"

	"github.com/deyvidm/sms-backend/models"
	"github.com/deyvidm/sms-backend/types"
	"github.com/gin-gonic/gin"
)

func AllEvents(c *gin.Context) {
	user, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}
	events, err := user.AllEvents()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": types.StatusSuccess, "data": events})
}

func NewEvent(c *gin.Context) {
	var input types.NewEventData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}

	user, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}

	err = user.SaveEvent(models.Event{
		Title: input.Ttile,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": types.StatusSuccess, "data": input})
}
