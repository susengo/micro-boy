package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/susengo/micro-boy/global"
	"github.com/susengo/micro-boy/model"
	"github.com/susengo/micro-boy/utils"
	"gorm.io/gorm"
)

// Login: 用户登录
// author: susen_go
// time: 2021-4-21
func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}

// Register: 用户注册
// author: susen_go
// time: 2021-04-19
func Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	// 判断用户名是否已注册
	if !errors.Is(global.GVA_DB.Where("username = ? ", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册"), userInter
	}

	// 密码md5简单加密，随机生成uuid，注册用户
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return err, u

}

// GetUserInfo: 通过uuid或username获取用户信息，测试用
// author: susen_go
// time: 2021-4-19
func GetUserInfo(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	// 判断用户名是否已注册
	err = global.GVA_DB.Where("uuid = ? OR username = ?", u.UUID, u.Username).First(&user).Error

	return err, user

}

// FindUserByUuid: 通过uuid获取用户信息
// author: susen_go
// time: 2021-4-20
func FindUserByUuid(uuid string) (err error, user *model.SysUser) {
	var u model.SysUser
	if err = global.GVA_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
