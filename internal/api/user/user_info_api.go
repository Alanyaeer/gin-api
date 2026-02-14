package user

import (
	"chat-system/internal/model/user"
	"chat-system/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

// # 1. 获取用户信息 (GET)
// curl -X GET "http://localhost:8888/api/v1/userInfo/getUserInfoByUserId?userId=42"

func GetUserInfoByUserId(ctx *gin.Context) {
	userId := ctx.Query("userId")
	fmt.Printf("userId type %T value %v\n", userId, userId)
	userInfo := service.GetUserInfoByUserId(userId)
	ctx.JSON(200, gin.H{
		"message": "get user info by user id",
		"data":    userInfo,
		"code":    200,
	})
}

// # 2. 添加用户信息 (POST)
//
//	curl -X POST "http://localhost:8888/api/v1/userInfo/addUserInfo" \
//	  -H "Content-Type: application/json" \
//	  -d '{
//	    "userId": "U2026214",
//	    "name": "张三",
//	    "age": 28,
//	    "sex": "男"
//	  }'
func AddUserInfo(ctx *gin.Context) {
	var userInfo user.UserInfo
	if err := ctx.ShouldBindJSON(&userInfo); err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid request body",
			"data":    0,
			"code":    400,
		})
		return
	} else {
		if err := service.AddUserInfo(userInfo); err != nil {
			ctx.JSON(500, gin.H{
				"message": "create user failed: " + err.Error(),
				"data":    nil,
				"code":    500,
			})
			return
		}
		ctx.JSON(200, gin.H{
			"message": "写入成功",
			"data":    nil,
			"code":    200,
		})
	}
}
