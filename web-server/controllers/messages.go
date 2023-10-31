package controllers

import (
	"net/http"

	"github.com/deyvidm/sms/dispatcher/tasks"
	"github.com/deyvidm/sms/web-server/types"
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

	contactNumber, err := GetContactNumber(c, input.To)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}

	t, err := tasks.NewMesssageTask(contactNumber, input.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err})
		return
	}
	ac := GetAsynqClient()
	_, err = ac.Enqueue(t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err})
		return
	}

	if err := user.SaveMessage(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": types.StatusSuccess})
}
