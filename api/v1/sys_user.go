package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
	"github.com/susengo/micro-boy/global"
	"github.com/susengo/micro-boy/middleware"
	"github.com/susengo/micro-boy/model"
	"github.com/susengo/micro-boy/model/request"
	"github.com/susengo/micro-boy/model/response"
	"github.com/susengo/micro-boy/service"
	"github.com/susengo/micro-boy/utils"
	"go.uber.org/zap"
	"time"
)

// Login: 用户登录
// 输入用户名、密码、验证码进行登录
func Login(c *gin.Context) {
	var L request.Login
	_ = c.ShouldBindJSON(&L)
	if err := utils.Verify(L, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if store.Verify(L.CaptchaId, L.Captcha, true) {
		U := &model.SysUser{Username: L.Username, Password: L.Password}
		if err, user := service.Login(U); err != nil {
			global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误", zap.Any("err", err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
		} else {
			tokenNext(c, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}
}

// tokenNext: 登录以后签发jwt
func tokenNext(c *gin.Context, user model.SysUser) {
	j := &middleware.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		NickName: user.NickName,
		Username: user.Username,
		//AuthorityId: user.AuthorityId,
		BufferTime: global.GVA_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "qmPlus",                                              // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	if err, jwtStr := service.GetRedisJWT(user.Username); err == redis.Nil {
		if err := service.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("设置登录状态失败", zap.Any("err", err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败", zap.Any("err", err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := service.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := service.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// Register: 用户注册
// 输入用户名、密码、昵称和头像进行注册，若昵称和头像为空，再自动填充默认值
func Register(c *gin.Context) {
	var R request.Register
	_ = c.ShouldBindJSON(&R)
	if err := utils.Verify(R, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	user := &model.SysUser{
		Username:  R.Username,
		Password:  R.Password,
		NickName:  R.NickName,
		HeaderImg: R.HeaderImg,
	}

	err, userReturn := service.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("注册失败", zap.Any("err", err))
		response.FailWithDetailed(response.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		response.OkWithDetailed(response.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

// TODO 测试用的，后续删掉
// GetUserInfo: 获取用户信息（通过username或uuid）
func GetUserInfo(c *gin.Context) {
	var user model.SysUser
	var err error

	username := c.Query("username")
	user.UUID, err = uuid.FromString(c.Query("uuid"))
	if err != nil && username == "" {
		response.FailWithMessage("格式不正确", c)
		return
	}

	user.Username = username
	err, userRe := service.GetUserInfo(user)
	if err != nil {
		response.FailWithMessage("获取用户信息失败", c)
		return
	} else {
		response.OkWithData(userRe, c)
	}
}
