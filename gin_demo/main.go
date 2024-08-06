package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

type OrderGetParam struct {
	OrderId   int32  `form:"orderId" json:"orderId"`
	OrderName string `form:"orderName" json:"orderName"`
}

// MyHandler 中间件 (Go中称为中间件，可以理解为 拦截器)
func MyHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 通过自定义的 中间件 也即 拦截器，后续处理时，只需要加载了中间件的，都可以获得到 这里的参数
		context.Set("userSession", 112233445566)
	}
}

func main() {

	// 请求地址
	ginServer := gin.Default()
	ginServer.Use(favicon.New("./resources/favicon.ico"))

	// 请求 接口
	ginServer.GET("/hi", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "Hello World"})
	})

	ginServer.POST("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "POST"})
	})
	ginServer.PUT("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "PUT"})
	})
	ginServer.DELETE("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "DELETE"})
	})

	// 加载静态页面
	ginServer.LoadHTMLGlob("templates/*")

	// 加载静态资源
	ginServer.Static("/resources/static", "./resources/static")

	// 返回html页面给前端
	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{"msg": "爽"})
	})

	/**
	参数接收
	*/

	// 普通的 Query 传参/order?orderId=1&orderName=MainOrder, 中间的参数 MyHandler 表示 中间件或拦截器，演示 拦截器中获取参数
	ginServer.GET("/order", MyHandler(), func(context *gin.Context) {
		/*orderId := context.Query("orderId")
		orderName := context.Query("orderName")
		context.JSON(http.StatusOK, gin.H{
			"orderId":   orderId,
			"orderName": orderName,
		})*/

		// 获取 中间件 中填充的值 并 转型为 string 类型
		userSession := context.MustGet("userSession").(int)

		var reqParam OrderGetParam
		err := context.BindQuery(&reqParam)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"msg": "Param cannot bind"})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"orderId":     reqParam.OrderId,
				"orderName":   reqParam.OrderName,
				"userSession": userSession,
			})
		}
	})

	// Path 方式传参
	ginServer.GET("/order/:orderId/:orderName", func(context *gin.Context) {
		orderId := context.Param("orderId")
		orderName := context.Param("orderName")
		context.JSON(http.StatusOK, gin.H{
			"orderId":   orderId,
			"orderName": orderName,
		})
	})

	// JSON 传参
	ginServer.POST("/order", func(context *gin.Context) {
		data, _ := context.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(data, &m)
		context.JSON(http.StatusOK, m)
	})

	/**
	路由
	*/

	// 301 重定向
	ginServer.GET("/redirect", func(context *gin.Context) {
		// 加上 https 或 http 表示 重定向到外部, /xx/yy 表示重定向到内部
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	// 404 跳转, 上面 加载过静态页面目录了，此处直接配置相对路径名称即可
	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	/**
	路由组, 统一管理 路由信息，省略了 func(context *gin.Context) {} 函数入参
	*/
	userGroup := ginServer.Group("/menu")
	{
		userGroup.GET("/info", func(context *gin.Context) {
			context.JSON(http.StatusOK, "hello menu info")
		})
		userGroup.POST("/create")
		userGroup.PUT("/update")
	}

	/**
	拦截器 MyHandler, 应用在 ginServer.GET("/order", MyHandler(), func(context *gin.Context) {}）
	*/

	// 启动服务
	_ = ginServer.Run(":8890")
}
