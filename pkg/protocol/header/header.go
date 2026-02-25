package header

import "github.com/gin-gonic/gin"

// 设置响应头为text/event-stream
func WrapperCtxHeaderForSse(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
}
