package main

import (
	_ "chat-system/pkg/log"
	"chat-system/internal/router"
)
func main() {
	e := router.SetupRouter()
	e.Run(":8888")
}