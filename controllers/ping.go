package controllers

import (
	"net/http"

	"github.com/deyvidm/sms-backend/types"
	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": types.StatusSuccess, "data": "pong"})
}
