package service

import (
	"errors"
	"github.com/susengo/micro-boy/global"
	"github.com/susengo/micro-boy/model"
	"gorm.io/gorm"
	"time"
)

// JsonInBlacklist: 拉黑jwt，加入jwt黑名单
func JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	err = global.GVA_DB.Create(&jwtList).Error
	return
}

// IsBlacklist: 判断jwt是否在黑名单内部
func IsBlacklist(jwt string) bool {
	isNotFound := errors.Is(global.GVA_DB.Where("jwt = ?", jwt).First(&model.JwtBlacklist{}).Error, gorm.ErrRecordNotFound)
	return !isNotFound
}

// GetRedisJWT: 从redis取jwt
func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.GVA_REDIS.Get(userName).Result()
	return err, redisJWT
}

// SetRedisJWT: jwt存入redis并设置过期时间
func SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.GVA_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.GVA_REDIS.Set(userName, jwt, timer).Err()
	return err
}
