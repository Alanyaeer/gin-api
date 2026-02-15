package response

import (
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, data any) {
	ctx.JSON(200, gin.H{
		"message": "success", 
		"data": data,
		"code": 200,
	})
}


func Error(ctx *gin.Context, message string, code int) {
	ctx.JSON(
		code, gin.H{
			"message": message,
			"data": nil,
			"code": code,
		},		
	)
}