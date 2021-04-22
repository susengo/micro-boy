package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/susengo/micro-boy/global"
	"github.com/susengo/micro-boy/model"
	"github.com/susengo/micro-boy/model/response"
	"github.com/susengo/micro-boy/service"
	"go.uber.org/zap"
)

// JsonInBlacklist 将jwt加入黑名单
func JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := model.JwtBlacklist{Jwt: token}
	if err := service.JsonInBlacklist(jwt); err != nil {
		global.GVA_LOG.Error("jwt作废失败!", zap.Any("err", err))
		response.FailWithMessage("jwt作废失败", c)
	} else {
		response.OkWithMessage("jwt作废成功", c)
	}
}
