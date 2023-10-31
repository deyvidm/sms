package controllers

import (
	"net/http"

	"github.com/deyvidm/sms/cmd/web-server/types"
	"github.com/gin-gonic/gin"
)

func EventDetails(c *gin.Context) {
	u, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}
	event := u.GetEventByID(c.Param("id"))
	// event, err := models.GetEventByID(c.Param("id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{"status": types.StatusSuccess, "data": event.ToAPI()})
}

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
	var input types.NewEvent

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}

	user, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}

	err = user.OrganizeEvent(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": types.StatusSuccess, "data": input})
}
