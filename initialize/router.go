package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/susengo/micro-boy/global"
	"github.com/susengo/micro-boy/middleware"
	"github.com/susengo/micro-boy/router"
	"net/http"
)

func Routers() *gin.Engine {
	// gin引擎
	var Router = gin.Default()
	// 提供文件存放在本地的静态地址
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path))

	// 跨域
	Router.Use(middleware.Cors())
	global.GVA_LOG.Info("micro-boy cors openning...")

	// 方便统一添加路由组前缀，多服务器上线使用
	PublicGroup := Router.Group("/public/v1")
	{
		// 注册基础功能路由，不做鉴权
		router.InitBaseRouter(PublicGroup)
	}
	PrivateGroup := Router.Group("/private/v1")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		// 注册用户路由
		router.InitUserRouter(PrivateGroup)
		// 注册jwt相关路由
		router.InitJwtRouter(PrivateGroup)
	}

	global.GVA_LOG.Info("micro-boy router register success")
	return Router

}
