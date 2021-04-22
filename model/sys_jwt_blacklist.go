package model

import (
	"github.com/susengo/micro-boy/global"
)

// jwt黑名单表
type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
