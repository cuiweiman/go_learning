package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {

	// 请求地址
	ginServer := gin.Default()
	ginServer.Use(favicon.New("./resources/favicon.ico"))

	// 请求 接口
	ginServer.GET("/hi", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "Hello World"})
	})

	// 启动服务
	ginServer.Run(":8890")
}
