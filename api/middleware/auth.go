package middleware

import (
	"github.com/Octek/resource-profile-management-backend.git/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ContextKey string

const (
	ContextUserKey ContextKey = "user"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		StaticToken := utils.GetBearerTokenString()
		tokenParts := strings.Split(authHeader, "Bearer ")
		if len(tokenParts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization format"})
			return
		}

		token := tokenParts[1]
		if token != StaticToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set(string(ContextUserKey), token)
		c.Next()
	}
}
