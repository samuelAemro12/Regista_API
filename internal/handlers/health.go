package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health returns the liveness status of the API.
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
