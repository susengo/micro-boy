package response

import (
	"github.com/susengo/micro-boy/model"
)

// 用户返回值结构体
type SysUserResponse struct {
	User model.SysUser `json:"user"`
}

// 登录成功返回的结构体
type LoginResponse struct {
	User      model.SysUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}
