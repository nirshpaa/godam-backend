package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORS middleware
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// During development, allow specific origins
		allowedOrigins := []string{
			"http://localhost:8081",
			"http://127.0.0.1:8081",
			"http://localhost:19006",
			"http://192.168.1.93:8081",
			"http://192.168.1.93:8000",
			"http://192.168.86.180:8081",
			"http://192.168.86.180:8000",
			"*", // Allow all origins for development
		}

		origin := c.Request.Header.Get("Origin")
		allowed := false

		// Check if the origin is in the allowed list
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin || allowedOrigin == "*" {
				allowed = true
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		// If not allowed, use the first origin as default
		if !allowed && len(allowedOrigins) > 0 {
			c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigins[0])
		}

		// Set CORS headers
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Company-ID")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // 24 hours
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Range")

		// Add additional headers for image handling
		c.Writer.Header().Set("Cross-Origin-Resource-Policy", "cross-origin")
		c.Writer.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		c.Writer.Header().Set("Cross-Origin-Opener-Policy", "same-origin")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
