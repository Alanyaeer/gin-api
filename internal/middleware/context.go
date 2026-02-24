package middleware

import (
	"github.com/gin-gonic/gin"
)

type contextKey string

const (
	RequestIDKey contextKey = "request_id" // 请求ID
)

// SetRequestID 设置请求ID
func SetRequestID(c *gin.Context, requestID string) {
	c.Set(RequestIDKey, requestID)
}

// GetRequestID 获取请求ID
func GetRequestID(c *gin.Context) string {
	if requestID, exists := c.Get(RequestIDKey); exists {
		if requestIDStr, ok := requestID.(string); ok {
			return requestIDStr
		}
	}
	return ""
}
