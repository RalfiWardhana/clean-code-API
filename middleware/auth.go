package middleware

import (
	"context"
	"errors"
	"strings"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gin-gonic/gin"
)

var PrivateKey = []byte("SuperSecretKey")

func DecryptJWT(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("auth invalid")
		}
		return PrivateKey, nil
	})
	if err != nil {
		return map[string]interface{}{}, err
	}

	if !parsedToken.Valid {
		return map[string]interface{}{}, err
	}

	return parsedToken.Claims.(jwt.MapClaims), nil
}

func WithAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
			c.JSON(401, map[string]string{
				"message": "unauthorize",
			})
			c.Abort()
		}

		auth := strings.Split(authHeader, " ")
		userData, err := DecryptJWT(auth[1])
		if err != nil {
			c.JSON(401, map[string]string{
				"message": "unauthorize",
			})
			c.Abort()
		}
		ctxUserID := context.WithValue(c.Request.Context(), "user_id", userData["id"])
		c.Request = c.Request.WithContext(ctxUserID)
		c.Next()
	}
}
