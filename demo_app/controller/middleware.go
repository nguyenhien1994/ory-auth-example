package controller

import (
	"github.com/gin-gonic/gin"
)

func AuthenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: parse jwts and set to context
		c.Next()
	}
}

// Authorize determines if current subject has been authorized to take an action on an object.
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: additional check
		c.Next()
	}
}
