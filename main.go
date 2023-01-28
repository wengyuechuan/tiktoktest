package main

import (
	"log"
	"tiktoktest/router"
)

func main() {
	e := router.Router()
	err := e.Run(":8080")
	if err != nil {
		log.Printf("服务器启动失败")
		return
	}
}
