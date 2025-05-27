package server

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func isAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.Info("Checking authentication")
		// Implement your authentication logic here
		// For example, check for a valid token in the request header
		token := c.GetHeader("Authorization")
		if token == "" {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
