package entity

import (
	"ginchat/utils"
	"gorm.io/gorm"
	"time"
)

// UserBasic 用户基本信息
type UserBasic struct {
	gorm.Model
	Identity      string
	Username      string
	Password      string
	Phone         string
	Email         string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time
	HeartbeatTime time.Time
	OfflineTime   time.Time
	IsLogout      bool
	DeviceInfo    string
}

// TableName mysql 表名
func (table *UserBasic) TableName() string {
	return "user_basic"
}

func ListUserBasic() []*UserBasic {
	var list = make([]*UserBasic, 10)
	utils.DB.Find(&list)
	return list
}

func CreateUser(basic *UserBasic) *gorm.DB {
	return utils.DB.Create(basic)
}

func GetUser(id int) *UserBasic {
	var user UserBasic = UserBasic{}
	// 先查询第一条记录 保存在 user 变量中, 等价于 select where id = #{id} limit 1
	utils.DB.Where("id = ?", id).First(&user)
	return &user
}

func DeleteUser(id int) bool {
	var user UserBasic = UserBasic{}
	// 先查询第一条记录 保存在 user 变量中, 等价于 select where id = #{id} limit 1
	utils.DB.Where("id = ?", id).First(&user)
	// 执行删除 delete from where id = 2
	utils.DB.Delete(&user)
	return true
}
