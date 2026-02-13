package router

import (
	"chat-system/internal/router/user"
	"github.com/gin-gonic/gin"
	"chat-system/config"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	apiGroup := r.Group(config.GlobalPrefixV1)
	user.RegisterUserInfoRoutes(apiGroup)
	return r
}