package main

import (
	"fmt"
	"github.com/taoti888/user/core"
	"github.com/taoti888/user/global"
	"github.com/taoti888/user/initialize"
	"github.com/taoti888/user/model"
	"gorm.io/gorm"
)

func RegisterTables(db *gorm.DB, tables ...any) {
	err := db.AutoMigrate(
		// 待注册的表
		tables...,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("register table success")
}

func main() {
	global.VP = core.Viper()

	// gorm连接数据库,注意程序结束前关闭数据库链接
	global.DB = initialize.Gorm()
	db, _ := global.DB.DB()
	defer db.Close()

	RegisterTables(
		global.DB,
		&model.Permissions{},
		&model.Role{},
		&model.User{},
		&model.UserRole{},
		&model.RolePermissions{},
	)
}
