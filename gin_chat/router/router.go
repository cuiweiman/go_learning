package router

import (
	_ "ginchat/docs"
	"ginchat/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	engine := gin.Default()

	// 为 swagger 注册路由
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	engine.GET("/index", service.GetIndex)
	engine.GET("/user/list", service.ListUserBasic)
	engine.POST("/user", service.CreateUserBasic)
	engine.DELETE("/user", service.DeleteUser)

	return engine
}
