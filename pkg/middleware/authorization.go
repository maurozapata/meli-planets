package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Authentication -
func Authentication(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.GetHeader("Authorization")

		if t == "" || t != token {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
			return
		}

		c.Next()
	}
}
