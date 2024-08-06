package test

import (
	"fmt"
	"ginchat/models/entity"
	"ginchat/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

func setup() {
	viper.SetConfigName("app")
	viper.AddConfigPath("../../conf")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	utils.InitMysql()
}

// TestListUserBasic 测试 获取 UserBasic 列表
func TestListUserBasic(t *testing.T) {
	setup()

	userList := entity.ListUserBasic()
	for i, userBasic := range userList {
		fmt.Println(i, userBasic)
	}
}

func TestCreateUser(t *testing.T) {
	setup()

	// 初始化一个 测试用的 用户信息
	userTest := &entity.UserBasic{
		Identity:      "test-001",
		Username:      "TestUser",
		Password:      "TestPassword",
		Phone:         "18877776666",
		Email:         "TestUser@ha.com",
		LoginTime:     time.Now(),
		HeartbeatTime: time.Now(),
		OfflineTime:   time.Now(),
	}
	// 创建测试用户
	entity.CreateUser(userTest)
	fmt.Println(userTest)
}

func TestDaoUserBasic(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:Baidu@123@tcp(127.0.0.1:3306)/gin_chat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema 当 实体类表 不存在时 自动创建
	db.AutoMigrate(&entity.UserBasic{})

	// 查询测试用户
	var queryUser entity.UserBasic
	// 根据 电话号码 字段查询
	db.First(&queryUser, "Phone = ?", "18877776666")
	log.Println(queryUser)

	// 更新
	//db.Model(&queryUser).Update("Email", "test-001@update.com")
	// 更新 多个字段
	//db.Model(&queryUser).Updates(models.UserBasic{Password: "UpdateMultiPassword", Phone: "18833330000"})
	//db.Model(&queryUser).Updates(map[string]interface{}{"Username": "MapUserName", "Email": "MapUserName@ha.com"})

	// 根据ID 删除
	//db.Delete(&queryUser, "test-001")
}
