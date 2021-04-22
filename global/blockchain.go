package global

import (
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/susengo/micro-boy/contract/helloworld"
)

// 区块链全局变量
var (
	// FISCO BCOS
	configsFisco      = []conf.Config{}                 // FISCO BCOS的配置信息
	helloworldSession = &helloworld.HelloWorldSession{} // FISCO BCOS的session句柄
)
