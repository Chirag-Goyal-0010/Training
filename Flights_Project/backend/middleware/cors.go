package middleware

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware handles CORS configuration
func CORSMiddleware() gin.HandlerFunc {
	// Get allowed origins from environment variable or use default
	allowedOrigins := getAllowedOrigins()

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Check if the origin is allowed
		if isAllowedOrigin(origin, allowedOrigins) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		// Set CORS headers
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // 24 hours

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// getAllowedOrigins returns the list of allowed origins
func getAllowedOrigins() []string {
	// Get origins from environment variable
	origins := os.Getenv("ALLOWED_ORIGINS")
	if origins == "" {
		// Default to localhost for development
		return []string{
			"http://localhost:3000", // React development server
			"http://localhost:8080", // Backend development server
			"http://127.0.0.1:3000",
			"http://127.0.0.1:8080",
		}
	}

	// Split the comma-separated list of origins
	return strings.Split(origins, ",")
}

// isAllowedOrigin checks if the given origin is in the allowed list
func isAllowedOrigin(origin string, allowedOrigins []string) bool {
	for _, allowed := range allowedOrigins {
		if origin == allowed {
			return true
		}
	}
	return false
}
