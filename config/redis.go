package config

// Redis的配置项
type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // 默认连接的数据库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // redis服务地址
	Password string `mapstructure:"password" json:"password" yaml:"password"` // redis访问密码
}
