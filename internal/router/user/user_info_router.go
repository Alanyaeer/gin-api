package user

import (
	"chat-system/internal/api/user"
	"github.com/gin-gonic/gin"
)

func RegisterUserInfoRoutes(r *gin.RouterGroup) {
	userInfoApiGroup := r.Group("/userInfo")
	{
		userInfoApiGroup.GET("/getUserInfoByUserId", user.GetUserInfoByUserId)
	}
}