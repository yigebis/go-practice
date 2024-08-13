package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var Users = make(map[string]User)

var secret_token = "123456"
var jwt_secret_signer = []byte(secret_token)

func GetMainPage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Welcome to the authentication page"})
}

func Register(c *gin.Context) {
	var user User

	bind_err := c.BindJSON(&user)
	if bind_err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid request payload"})
	}

	// Hashing and storing
	hashedPassword, hash_err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if hash_err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	}

	user.Password = string(hashedPassword)
	Users[user.Email] = user
}

func login(c *gin.Context) {
	var user User

	bind_err := c.BindJSON(&user)
	if bind_err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid request payload"})
	}

	//check if email exists
	existingUser, ok := Users[user.Email]
	if !ok {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	//check password match
	if bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	//grant jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": existingUser.ID,
		"email":   existingUser.Email,
	})

	jwtToken, err := token.SignedString(jwt_secret_signer)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "User logged in successfully", "token": jwtToken})

}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement JWT validation logic
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return jwt_secret_signer, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func Secure(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Welcome to the dashboard"})
}

func main() {
	router := gin.Default()

	router.GET("/", GetMainPage)
	router.POST("/register", Register)
	router.POST("/login", login)
	router.GET("/secure", Secure)
	router.Run()
}
