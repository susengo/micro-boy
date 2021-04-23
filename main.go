package main

import (
	"github.com/susengo/micro-boy/core"
	"github.com/susengo/micro-boy/global"
	"github.com/susengo/micro-boy/initialize"
)

func main() {
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = initialize.Gorm()
	if global.GVA_DB != nil {
		initialize.MysqlTables(global.GVA_DB)
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

	/*
		区块链初始化
	*/

	core.RunWindowsServer()
}
