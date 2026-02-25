package main

import (
	"chat-system/config"
	"chat-system/internal/router"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	config.InitConfig(".", "config", "yaml")
	port := ":" + strconv.Itoa(config.Cfg.App.Port)
	router.SetupRouter(e)
	e.Run(port)
}
