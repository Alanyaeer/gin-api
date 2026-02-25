package user

import (
	"chat-system/internal/api/user"
	"github.com/gin-gonic/gin"
)

type UserRouteRegistrar struct{}

func NewRouter() *UserRouteRegistrar {
	return &UserRouteRegistrar{}
}

func (u *UserRouteRegistrar) RegisterRoutes(r *gin.RouterGroup) {
	userInfoAPIGroup := r.Group("/userInfo")
	{
		userInfoAPIGroup.GET("/getUserInfoByUserId", user.GetUserInfoByUserId)
		userInfoAPIGroup.POST("/addUserInfo", user.AddUserInfo)
	}
}