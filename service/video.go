package service

import "github.com/gin-gonic/gin"

func GetVideo(c *gin.Context) {
	path := "D:/goproject/tiktoktest/video/"
	// 解析请求视频的地址
	fileName := path + c.Query("name")

	// 返回一个视频
	c.File(fileName)
}
