package main

import (
	"github.com/taoti888/user/core"
	"github.com/taoti888/user/global"
	"github.com/taoti888/user/initialize"
	"go.uber.org/zap"
)

func main() {
	// 初始化Viper
	global.VP = core.Viper()

	// 初始化zap日志库
	global.LOG = core.Zap()
	zap.ReplaceGlobals(global.LOG)

	// gorm连接数据库,注意程序结束前关闭数据库链接
	global.DB = initialize.Gorm()
	db, _ := global.DB.DB()
	defer db.Close()

	// 运行程序
	core.RunWindowsServer()
}
