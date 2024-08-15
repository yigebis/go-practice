package controllers

import (
	"net/http"
	"task_management_api_with_clean_architecture/domain"
	"task_management_api_with_clean_architecture/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase *usecase.UserUsecase
}

func NewUserController(uuc *usecase.UserUsecase) *UserController {
	return &UserController{UserUseCase: uuc}
}

func (uc *UserController) Register(c *gin.Context) {
	var user domain.User

	bind_err := c.BindJSON(&user)
	if bind_err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid request payload"})
		return
	}

	if user.Username == "" || user.Password == "" {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "username and password can't be empty"})
		return
	}

	code, err := uc.UserUseCase.RegisterUser(user)

	if err != nil {
		c.IndentedJSON(code, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "user has been registered successfully"})
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (uc *UserController) Login(c *gin.Context) {
	var credential Credentials

	bind_err := c.BindJSON(&credential)
	if bind_err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid request payload"})
		return
	}

	token, code, err := uc.UserUseCase.LogUser(credential.Username, credential.Password)

	if err != nil {
		c.IndentedJSON(code, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})

}

func (uc *UserController) Promote(c *gin.Context) {
	username := c.Param("username")

	code, err := uc.UserUseCase.Promote(username)

	if err != nil {
		c.IndentedJSON(code, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User promoted successfully"})
}
