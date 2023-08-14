package gen

const TempBoot = `
package boot

import (
	"TempImportPkg/global"
	"TempImportPkg/infra/config"
	"TempImportPkg/infra/db/mysql"
)

// 初始化
func init() {
	global.GConfig = config.NewConfig() // 初始化全局配置
	global.GDB = mysql.NewDB()          // 初始化全局DB
}

`
