package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/susengo/micro-boy/api/v1"
)

// 用户相关路由
func InitUserRouter(Router *gin.RouterGroup) {
	// 配置路由和中间件
	UserRouter := Router.Group("user").Use()
	{
		// 获取用户信息
		UserRouter.GET("info", v1.GetUserInfo)
	}

}
