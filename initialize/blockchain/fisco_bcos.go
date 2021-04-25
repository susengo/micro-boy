package blockchain

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/susengo/micro-boy/contract/helloworld"
	"github.com/susengo/micro-boy/global"
	"go.uber.org/zap"
)

/*
	TODO: 这个地方需要优化
	部署FISCO BCOS智能合约，获取合约调用地址，再初始化合约调用句柄，后续合约升级迭代怎么做？
*/

var (
	configFisco       = []conf.Config{}
	helloworldSession = &helloworld.HelloWorldSession{}
)

// 读取FISCO BCOS配置信息，安装智能合约代码
// 返回合约地址及合约部署交易哈希
func InitFiscoDeploy() {
	var err error
	// 读取配置参数
	configFisco, err = conf.ParseConfigFile("config.toml")
	if err != nil {
		global.GVA_LOG.Error("ParseFISCOConfigFile failed.", zap.Any("err", err))
	}

	config := &configFisco[0]
	clientFisco, err := client.Dial(config)
	if err != nil {
		global.GVA_LOG.Fatal("InitFiscoDeploy: 连接FISCO BCOS失败.", zap.Any("err", err))
	}

	// 安装合约
	address, tx, instance, err := helloworld.DeployHelloWorld(clientFisco.GetTransactOpts(), clientFisco)
	if err != nil {
		global.GVA_LOG.Fatal("FISCO BCOS部署智能合约失败.", zap.Any("err", err))
	}

	// addressString: 智能合约部署成功的调用地址
	// txString: 合约部署交易哈希
	addressString := address.Hex()
	txString := tx.Hash().Hex()
	_ = instance
	// 将configFisco赋值给全局变量
	global.ConfigsFisco = configFisco
	global.GVA_LOG.Info("Helloworld contract address:", zap.Any("addressString", addressString))
	global.GVA_LOG.Info("Helloworld transaction hash:", zap.Any("txString", txString))
	global.GVA_LOG.Info("FISCO BCOS全局变量为:", zap.Any("FISCO Config", fmt.Sprintf("%+v", global.ConfigsFisco)))

}

// 实例化合约调用，传入合约调用地址
// helloworld合约测试地址：0xcB7EBC148292F17BcA3aF0E44d0eC80A56b21A5b
func InitFiscoSession(address string) {
	var err error
	// 读取配置参数
	configFisco, err = conf.ParseConfigFile("config.toml")
	if err != nil {
		global.GVA_LOG.Error("ParseFISCOConfigFile failed,err: %v", zap.Any("err", err))
	}

	config := &configFisco[0]
	clientFisco, err := client.Dial(config)
	if err != nil {
		global.GVA_LOG.Fatal("InitFiscoSession: 连接FISCO BCOS失败.", zap.Any("err", err))
	}

	contractAddress := common.HexToAddress(address)
	instance, err := helloworld.NewHelloWorld(contractAddress, clientFisco)
	if err != nil {
		global.GVA_LOG.Error("err", zap.Any("err", err))
	}

	helloworldSession = &helloworld.HelloWorldSession{Contract: instance, CallOpts: *clientFisco.GetCallOpts(), TransactOpts: *clientFisco.GetTransactOpts()}
	// 将helloworldSession赋值给全局变量
	global.HelloworldSession = helloworldSession
	global.GVA_LOG.Info("##### 初始化智能合约Session成功 #####")
	global.GVA_LOG.Info("FISCO BCOS session为:", zap.Any("FISCO Config", fmt.Sprintf("%+v", global.HelloworldSession)))
}
