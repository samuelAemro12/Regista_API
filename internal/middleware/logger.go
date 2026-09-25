package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger records basic request timing and response status information.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		started := time.Now()
		c.Next()
		log.Printf("%s %s %d %s", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), time.Since(started))
	}
}
