package global

import (
	"go.uber.org/zap"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"github.com/susengo/micro-boy/config"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB      // 数据库对象
	GVA_REDIS  *redis.Client // Redis对象
	GVA_CONFIG config.Server // 服务总配置对象
	GVA_VP     *viper.Viper  // 配置库解析对象
	GVA_LOG    *zap.Logger   // 日志对象
)
