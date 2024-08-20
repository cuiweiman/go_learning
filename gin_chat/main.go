package main

import (
	"fmt"
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

	serverInfo := `
	Url: http://127.0.0.1:8081
	Swagger: http://localhost:8081/swagger/index.html

`
	fmt.Printf(serverInfo)
	// listen and serve on 0.0.0.0:8080
	engine.Run(":8081")
}
