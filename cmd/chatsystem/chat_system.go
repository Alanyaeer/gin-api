package main

import (
	"chat-system/config"
	"chat-system/internal/router"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.New()
	config.InitConfig(".", "config", "yaml")
	port := ":" + strconv.Itoa(config.Cfg.App.Port)
	router.SetupRouter(e)
	e.Run(port)
}
