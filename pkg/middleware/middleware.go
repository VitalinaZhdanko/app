// Package middleware implements middlewares for routers
package middleware

import (
	"app/diplom/pkg/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

)

// Auth checks if a valid token exists
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenAuth := c.Request.Header.Get("Authorization")
		id, err := token.Parse(strings.TrimPrefix(tokenAuth, "Bearer "))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid token")
			return
		}
		c.Set("user_id", id)
		c.Next()
	}
}
