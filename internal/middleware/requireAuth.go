package middleware

import (
	"github.com/Yelsnik/tracking-inventory.git/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RequireAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication token"})
			c.Abort()
			return
		}

		// The token should be prefixed with "Bearer "
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
			c.Abort()
			return
		}

		tokenString = tokenParts[1]

		claims, err := utils.Validate(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
			c.Abort()
			return
		}

		c.Set("sub", claims["sub"])

		c.Next()
	}
}
