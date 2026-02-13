package main

import (
	"chat-system/internal/router"
)
func main() {
	e := router.SetupRouter()
	e.Run(":8888")
}