package middleware

import (
	// "task_manager_with_auth/models"

	"fmt"
	"net/http"
	"strings"

	// "time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwt_signer = "123456" // Replace with a secure secret

// func GenerateJWT(user models.User) (string, error) {
//     token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//         "id":       user.ID,
//         "username": user.Username,
//         "role":     user.Role,
//         "exp":      time.Now().Add(time.Hour * 72).Unix(),
//     })

//     return token.SignedString(jwtSecret)
// }

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
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

			return []byte(jwt_signer), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Set("user", claims)
		fmt.Println(claims)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("user")
		if !exists {
			fmt.Printf("Claims: %v\n", claims)
			c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
			c.Abort()
			return
		}

		userClaims := claims.(jwt.MapClaims)
		role := userClaims["role"].(string)
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
