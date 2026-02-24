package response

import (
	"github.com/gin-gonic/gin"
)

type StatusCode int

const (
	ServerCode  StatusCode = 500
	ClientCode  StatusCode = 400
	SuccessCode StatusCode = 200
)

func Success(ctx *gin.Context, data any) {
	ctx.JSON(int(SuccessCode), gin.H{
		"message": "success",
		"data":    data,
		"code":    200,
	})
}

func Error(ctx *gin.Context, message string, statusCode StatusCode) {
	ctx.JSON(
		int(statusCode), gin.H{
			"message": message,
			"data":    nil,
			"code":    int(statusCode),
		},
	)
}
