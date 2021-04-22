package router

import (
	"github.com/gin-gonic/gin"
	"github.com/susengo/micro-boy/api/v1"
)

// jwt相关路由
func InitJwtRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("jwt").Use()
	{
		ApiRouter.POST("jsonInBlacklist", v1.JsonInBlacklist) // jwt加入黑名单
	}
}
