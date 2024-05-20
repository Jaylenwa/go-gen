package gen

const TempInit = `
package boot

import (
	"TempImportPkg/global"
	"TempImportPkg/init/config"
	"TempImportPkg/init/db"
	"TempImportPkg/init/log"
)

// 初始化
func init() {
	global.Config = config.InitConfig()
	global.DB = db.InitDB()
	global.Log = log.InitLog()
}
`
