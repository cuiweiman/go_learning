初始化创建项目

> - [go package](https://pkg.go.dev/)
> - [gorm docs](https://gorm.io/docs/)

1. 新建 gin_chat 空文件夹
2. 创建空白 main.go 文件
    ```go
    package main
    
    func main() {
    
    }
    ```
3. 初始化 mod 模块
    ```bash
    go mod init ginchat
    go mod tidy
    ```
4. 编码 models 模块
   ```bash
   # 当 gorm.Model 的包 拉不下来是，手动执行
   go get -u gorm.io/gorm
   go get -u github.com/go-gorm/gorm
   ```

5. 引入Gin
   ```bash
   # https://pkg.go.dev/github.com/gin-gonic/gin
   go get -u github.com/gin-gonic/gin 
   ```

6. 引入Swagger
   ```bash
   # https://github.com/swaggo/swag/blob/v2.0.0-rc3/README_zh-CN.md#%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B
   go install github.com/swaggo/swag/cmd/swag@latest
   
   # main.go 路径下执行,生成 docs 文档并需要导入到 router 中。 按照官方文档操作
   swag init 
   
   ```
   