package router

import (
	"chat-system/internal/router/user"
	"chat-system/internal/validator/userinfo" 
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"chat-system/config"

)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	// 解决 gin 框架的一个安全问题，详见 https://github.com/gin-gonic/gin/blob/master/docs/doc.md#dont-trust-all-proxies
	// 如果服务器不通过代理访问，为了消除启动的安全警告，设置为nil
	r.SetTrustedProxies(nil)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NameValid", userinfo.NameValid)
	}
	apiGroup := r.Group(config.GlobalPrefixV1)
	user.RegisterUserInfoRoutes(apiGroup)
	return r
}