package utils

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(out io.Writer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()
		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := func(c *gin.Context) string {
			if ips := c.GetHeader("x-forwarded-for"); ips != "" {
				ip := strings.Split(ips, ",")
				return ip[0]
			}
			if ip := c.GetHeader("Proxy-Client-IP"); ip != "" {
				return ip
			}
			if ip := c.GetHeader("X-Real-IP"); ip != "" {
				return ip
			}
			return c.ClientIP()
		}(c)
		method := c.Request.Method
		statusCode := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		fmt.Fprintf(out, "[GIN] %v [%3d] %13v | %15s |%-7s %s\n",
			end.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)

	}
}
