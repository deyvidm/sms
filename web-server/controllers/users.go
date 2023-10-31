package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/deyvidm/sms/web-server/auth"
	"github.com/deyvidm/sms/web-server/models"
	"github.com/deyvidm/sms/web-server/types"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input types.NewUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}

	u := models.User{
		Username: input.Username,
		Password: input.Password,
		Contact: models.Contact{
			FirstName: input.FirstName,
			LastName:  input.LastName,
		},
	}
	_, err := u.RegisterUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": "failed to save user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": types.StatusSuccess, "data": fmt.Sprintf("welcome %s!", input.Username)})
}

func Login(c *gin.Context) {
	var input types.LoginUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}

	u := models.User{
		Username: input.Username,
		Password: input.Password,
	}
	apiUser, token, err := models.LoginUser(u.Username, u.Password)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": "incorrect login details"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": types.StatusSuccess, "data": map[string]interface{}{"token": token, "user": apiUser}})
}

func GetUserFromContext(c *gin.Context) (models.User, error) {
	userID, err := auth.ExtractTokenID(c)
	if err != nil {
		return models.User{}, err
	}
	u, err := models.GetUserByID(userID)
	u.Password = "no :)"
	return u, err
}

func CurrentUser(c *gin.Context) {
	u, err := GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": types.StatusFailed, "data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": types.StatusSuccess, "data": u})
}

func GetContactNumber(c *gin.Context, id string) (string, error) {
	return models.GetContactNumber(id)
}
