package router

import (
	"github.com/gin-gonic/gin"
	"github.com/susengo/micro-boy/api/v1"
	"github.com/susengo/micro-boy/middleware"
)

// 基础路由
// 不需要经过鉴权的公共路由
func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base").Use(middleware.NeedInit())
	{
		// 用户注册
		BaseRouter.POST("register", v1.Register)
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("captcha", v1.Captcha)
	}
	return BaseRouter
}
