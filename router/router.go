package router

import (
	"github.com/gin-gonic/gin"
	"tiktoktest/service"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 构建一个简单的文件服务器
	// 用于测试返回内容
	r.GET("/api", service.GetJson)
	r.GET("/video", service.GetVideo)
	r.GET("/img", service.GetImg)
	return r
}
