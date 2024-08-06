package main

import (
	"ginchat/router"
	"ginchat/utils"
)

// @title Gin_Chat Demo API Docs
// @version 1.0
// @description 这是Golang开发的简单聊天Demo
// @host      localhost:8081
// @BasePath /
func main() {
	utils.InitConfig()
	utils.InitMysql()

	// 获取 Gorm 引擎
	engine := router.Router()
	// listen and serve on 0.0.0.0:8080
	engine.Run(":8081")
}
