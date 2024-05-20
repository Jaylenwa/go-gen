package gen

const TempGlobal = `
package global

import (
	"TempImportPkg/init/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config *config.Config
	Log    *logrus.Logger
)
`
