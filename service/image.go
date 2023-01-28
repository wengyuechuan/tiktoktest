package service

import "github.com/gin-gonic/gin"

func GetImg(c *gin.Context) {
	path := "D:/goproject/tiktoktest/img/"
	// 解析请求视频的地址
	fileName := path + c.Query("name")

	// 返回一个视频
	c.File(fileName)
}
