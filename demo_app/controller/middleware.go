package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

func AuthenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: parse jwts and set to context
		log.Println(c.Request.Header)
		c.Next()
	}
}

// Authorize determines if current subject has been authorized to take an action on an object.
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context	) {
		// TODO: additional check
		c.Next()
	}
}
