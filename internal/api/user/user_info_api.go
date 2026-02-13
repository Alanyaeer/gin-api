package user

import "github.com/gin-gonic/gin"
func GetUserInfoByUserId(ctx * gin.Context) (){
	ctx.JSON(200, gin.H{
		"message": "get user info by user id",
		"data": nil,
		"code": 0,
	})
}