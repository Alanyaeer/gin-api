package user

import (
	"chat-system/internal/model/dto"
	"chat-system/internal/service"
	"chat-system/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// # 1. 获取用户信息 (GET)
// curl -X GET "http://localhost:8888/api/v1/userInfo/getUserInfoByUserId?userId=42"

func GetUserInfoByUserId(ctx *gin.Context) {
	userId := ctx.Query("userId")
	fmt.Printf("userId type %T value %v\n", userId, userId)
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		response.Error(ctx, "invalid userId", response.CodeParamError)
		return
	}
	userInfo := service.GetUserInfoByUserId(userIdInt)
	response.Success(ctx, userInfo)
}

// # 2. 添加用户信息 (POST)
//
// curl -X POST "http://localhost:8888/api/v1/userInfo/addUserInfo" \
//   -H "Content-Type: application/json" \
//   -d '{
//     "userId": "U2026214",
//     "name": "张三",
//     "age": 28,
//     "sex": "男"
//   }'
func AddUserInfo(ctx *gin.Context) {
	var userInfoReq dto.UserInfoReq
	if err := ctx.ShouldBindJSON(&userInfoReq); err != nil {
		response.Error(ctx, "invalid request body", response.CodeParamError)
		return
	} else {
		if err := service.AddUserInfo(&userInfoReq); err != nil {
			response.Error(ctx, "invalid request body", response.CodeServerError)
			return
		}
		response.Success(ctx, userInfoReq)
	}
}
