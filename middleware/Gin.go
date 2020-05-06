package middleware

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"github.com/gin-contrib/cors"
	"os"
)

// GinCustomLogger returns custom logger for gin
func GinCustomLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s]: %s %s - %d(%s) [%s %s %s]",
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Request.UserAgent(),
			param.Request.Proto)
	})
}


// GinCors returns cors middleware for gin
func GinCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("REQUEST_URI")},
	})
}