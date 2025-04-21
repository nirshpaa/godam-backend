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
			"https://qck0g7e-anonymous-8081.exp.direct",
			"http://qck0g7e-anonymous-8081.exp.direct",
		}

		origin := c.Request.Header.Get("Origin")
		allowed := false

		// Check if the origin is in the allowed list
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				allowed = true
				c.Header("Access-Control-Allow-Origin", origin)
				break
			}
		}

		// If not allowed, use the first origin as default
		if !allowed && len(allowedOrigins) > 0 {
			c.Header("Access-Control-Allow-Origin", allowedOrigins[0])
		}

		// Set CORS headers
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Company-ID")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Header("Access-Control-Max-Age", "86400") // 24 hours
		c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Range")

		// Add additional headers for image handling
		c.Header("Cross-Origin-Resource-Policy", "cross-origin")
		c.Header("Cross-Origin-Embedder-Policy", "require-corp")
		c.Header("Cross-Origin-Opener-Policy", "same-origin")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
