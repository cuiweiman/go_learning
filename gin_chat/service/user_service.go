package service

import (
	"ginchat/models/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// ListUserBasic
// @Tags 用户管理
// @Summary 获取用户列表
// @Description 获取用户列表
// @Produce json
// @Success 200 {string} json{"code","data"}
// @Failure 500 {string} json{"code","data"}
// @Router /user/list [get]
func ListUserBasic(ctx *gin.Context) {
	userList := entity.ListUserBasic()
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "OK",
		"data": userList,
	})
}

// CreateUserBasic
// @Tags 用户管理
// @Summary 创建用户
// @Description 创建用户
// @Produce json
// @Param name query string false "用户名"
// @Param password query string false "密码"
// @Param rePassword query string false "确认密码"
// @Param phone query string false "手机号码"
// @Success 200 {string} json{"code","msg","data"}
// @Failure 500 {string} json{"code","msg","data"}
// @Router /user [post]
func CreateUserBasic(ctx *gin.Context) {
	var user entity.UserBasic = entity.UserBasic{
		Identity:      "test-001",
		Email:         "TestUser@ha.com",
		LoginTime:     time.Now(),
		HeartbeatTime: time.Now(),
		OfflineTime:   time.Now(),
	}
	user.Username = ctx.Query("name")
	user.Phone = ctx.Query("phone")
	user.Password = ctx.Query("password")
	rePassword := ctx.Query("rePassword")
	if user.Password != rePassword {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "密码不一致",
		})
	}
	entity.CreateUser(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "OK",
		"data": user.ID,
	})
}

// UpdateUserBasic
// @Tags 用户管理
// @Summary 更新手机号码和邮箱
// @Description 更新手机号码和邮箱
// @Produce json
// @Param id query int true "用户ID"
// @Param phone query string false "手机号码"
// @Param email query string false "电子邮箱"
// @Success 200 {string} json{"code","msg","data"}
// @Failure 500 {string} json{"code","msg","data"}
// @Router /user [put]
func UpdateUserBasic(ctx *gin.Context) {
	var idParam = ctx.Query("id")
	userId, _ := strconv.Atoi(idParam)
	userBasic := entity.GetUser(userId)
	if userBasic.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户信息不存在",
		})
	}
	userBasic.Phone = ctx.Query("phone")
	userBasic.Email = ctx.Query("email")

	entity.UpdateUser(userBasic)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "OK",
		"data": userBasic.ID,
	})
}

// DeleteUser
// @Tags 用户管理
// @Summary 删除用户
// @Description 删除用户
// @Produce json
// @Param id query int false "用户ID"
// @Success 200 {string} json{"code","msg","data"}
// @Failure 500 {string} json{"code","msg","data"}
// @Router /user [delete]
func DeleteUser(ctx *gin.Context) {
	var userId = ctx.Query("id")
	id, _ := strconv.Atoi(userId)
	getUser := entity.GetUser(id)
	if getUser.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "delete error, user not exists. UserId: " + userId,
		})
	}
	deleteUser := entity.DeleteUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "OK",
		"data": deleteUser,
	})
}
