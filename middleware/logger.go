package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger middleware
func Logger(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		param := gin.LogFormatterParams{
			Path:         path,
			Method:       c.Request.Method,
			StatusCode:   c.Writer.Status(),
			Latency:      time.Since(start),
			ClientIP:     c.ClientIP(),
			RawQuery:     raw,
			ErrorMessage: c.Errors.ByType(gin.ErrorTypePrivate).String(),
		}

		logger.Printf("[GIN] %s | %3d | %13v | %15s | %s | %s\n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
		)
	}
}
