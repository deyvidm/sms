package middleware

import (
	"net/http"

	"github.com/deyvidm/sms-backend/auth"
	"github.com/deyvidm/sms-backend/types"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": types.StatusFailed, "data": "you need to authenticate for this request"})
			c.Abort()
			return
		}
		c.Next()
	}
}
