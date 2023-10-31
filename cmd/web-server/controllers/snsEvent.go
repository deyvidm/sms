package controllers

import (
	"net/http"

	"github.com/deyvidm/sms/cmd/web-server/models"
	"github.com/deyvidm/sms/cmd/web-server/types"
	"github.com/deyvidm/sms/pkg/tasks"
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
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err})
		return
	}

	t, err := tasks.NewReponseTask(e.Message.OriginationNumber, e.Message.MessageBody)
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

	c.JSON(http.StatusOK, gin.H{"status": types.StatusSuccess, "data": e})
}
