package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/regista/regista-api/internal/handlers"
)

// NewRouter creates the API router. Middleware is registered by application
// wiring so the router remains focused on route definitions.
func NewRouter() *gin.Engine {
	return gin.New()
}

// RegisterAPI registers all versioned API routes.
func RegisterAPI(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/health", handlers.Health)
}
