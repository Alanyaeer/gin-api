package main

import (
	"chat-system/config"
	"chat-system/internal/router"
	_ "chat-system/pkg/log"
	"strconv"
)
func main() {
	e := router.SetupRouter()
	config.InitConfig(".", "config", "yaml")
	port := ":" + strconv.Itoa(config.Cfg.App.Port)
	e.Run(port)
}