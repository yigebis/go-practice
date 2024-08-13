package controllers

import (
	"net/http"
	"task_manager_with_auth/data"
	"task_manager_with_auth/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User

	bind_err := c.BindJSON(&user)
	if bind_err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid request payload"})
		return
	}

	if user.Username == "" || user.Password == "" {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "username and password can't be empty"})
		return
	}

	// Hashing and replacing password
	code, err := data.RegisterUser(user)

	if err != nil {
		c.IndentedJSON(code, gin.H{"error": err.Error()})
		return
	}

}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var credential Credentials

	bind_err := c.BindJSON(&credential)
	if bind_err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid request payload"})
		return
	}

	token, code, err := data.LogUser(credential.Username, credential.Password)

	if err != nil {
		c.IndentedJSON(code, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"token": token})

}

func Promote(c *gin.Context) {
	username := c.Param("username")

	code, err := data.Promote(username)

	if err != nil {
		c.IndentedJSON(code, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User promoted successfully"})
}
