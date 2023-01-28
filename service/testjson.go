package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type HelloJson struct {
	List []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"list"`
}

func TestHelloJson(c *gin.Context) {
	filePtr, err := os.Open("json/hello.json")
	if err != nil {
		log.Printf("文件打开失败 [ERR : %v]", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "文件不存在",
		})
		return
	}
	defer filePtr.Close()
	var info HelloJson
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		log.Printf("解码失败 [ERR : %v]", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "文件解码失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  info,
	})
}

// GetJson 这里在之后构建数据库的时候还需要在这里进行查看，我们先按照固定的格式存储
func GetJson(c *gin.Context) {
	path := "D:/goproject/tiktoktest/json/"
	// 解析请求视频的地址
	fileName := path + c.Query("jsonname")

	// 返回一个视频
	c.File(fileName)
}
