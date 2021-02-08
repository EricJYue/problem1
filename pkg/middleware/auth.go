package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// security achieved by jwt
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, _ := c.GetQuery("token")
		if len(tokenString) == 0 {
			tokenString = c.GetHeader("Authorization")
		}
		if len(tokenString) == 0 {
			c.AbortWithStatusJSON(400, gin.H{"msg": "empty token."})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("validation error")
			}
			return []byte("abc"), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"msg": "auth failed."})
			return
		}
		if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
			c.AbortWithStatusJSON(400, gin.H{"msg": "auth failed."})
			return
		}

	}
}
