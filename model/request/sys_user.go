package request

import uuid "github.com/satori/go.uuid"

// 用户注册请求的结构体
type Register struct {
	Username  string `json:"userName"`
	Password  string `json:"passWord"`
	NickName  string `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	//AuthorityId string `json:"authorityId" gorm:"default:888"`
}

// 用户登录请求的结构体
type Login struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

// 修改密码请求的结构体
type ChangePasswordStruct struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// Modify  user's auth structure
type SetUserAuth struct {
	UUID uuid.UUID `json:"uuid"`
	//AuthorityId string    `json:"authorityId"`
}
