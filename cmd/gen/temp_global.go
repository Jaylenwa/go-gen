package gen

const TempGlobal = `
package global

import (
	"TempImportPkg/infra/config"
	"gorm.io/gorm"
)

var (
	GConfig *config.Config // 全局配置
	GDB     *gorm.DB       // 全局 DB
)


`
