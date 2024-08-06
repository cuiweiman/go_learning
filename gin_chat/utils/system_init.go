package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

func InitMysql() {
	newLogger := logger.New(
		// 输出位置
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		// 输出内容 级别 和 颜色
		logger.Config{
			// 慢 SQL
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	var mysqlUrl = viper.GetString("mysql.dns")
	DB, _ = gorm.Open(mysql.Open(mysqlUrl), &gorm.Config{
		Logger: newLogger,
	})
}
