package middleware

import (
	"chat-system/pkg/response"
	"fmt"
	"runtime/debug"

	"log/slog"

	"github.com/gin-gonic/gin"
)

// Recovery 中间件 - 恢复 panic
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Get request ID
				requestID := GetRequestID(c)
				// tenantID := GetTenantID(c)

				// Build stack trace
				stack := debug.Stack()

				// Log error
				slog.Error("Request panic",
					"request_id", requestID,
					// "tenant_id", tenantID,
					"path", c.Request.URL.Path,
					"method", c.Request.Method,
					"ip", c.ClientIP(),
					"error", err,
					"stack", string(stack),
				)

				// 返回500错误
				response.Error(c, fmt.Sprintf("%v", err), 500)
				c.Abort()
			}
		}()

		c.Next()
	}
}
