package gen

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type TableColumn struct {
	Field   string `gorm:"column:Field"`   // 字段名称
	Type    string `gorm:"column:Type"`    // 字段类型
	Null    string `gorm:"column:Null"`    // 是否空
	Key     string `gorm:"column:Key"`     // 索引
	Default string `gorm:"column:Default"` // 默认值
	Extra   string `gorm:"column:Extra"`   // 扩展
	Comment string `gorm:"column:Comment"` // 备注
}

type GenReq struct {
	TableName         string
	TableKey          string
	BaseDir           string
	InitDir           string
	EntityDir         string
	ServiceDir        string
	GlobalDir         string
	InfraDir          string
	ConfigDir         string
	EnumsDir          string
	DBDir             string
	LogDir            string
	RepositoryPoDir   string
	RepositoryImplDir string
	DtoDir            string
	middlewareDir     string
	DrivenDir         string
	DriverDir         string
	HandlerDir        string
	RouterDir         string
	ManifestDir       string
	I18nDir           string
	InitConfigDir     string
	RequestDir        string
	ResponseDir       string
	UtilsDir          string
	TableColumns      []TableColumn
}

func GenInit(tableName string) GenReq {
	codePath := viper.Get("server.code_path").(string)
	baseDir := codePath + "/" + viper.Get("server.go_module").(string)

	return GenReq{
		BaseDir:           baseDir,
		TableName:         tableName,
		TableKey:          viper.Get("mysql.key").(string),
		InitDir:           baseDir + "/init/",
		EntityDir:         baseDir + "/domain/entity/",
		ServiceDir:        baseDir + "/domain/service/",
		GlobalDir:         baseDir + "/global/",
		InfraDir:          baseDir + "/infra",
		InitConfigDir:     baseDir + "/init/config/",
		EnumsDir:          baseDir + "/domain/enums/",
		DBDir:             baseDir + "/init/db/",
		LogDir:            baseDir + "/init/log/",
		RepositoryPoDir:   baseDir + "/infra/po/",
		RepositoryImplDir: baseDir + "/adapter/driven/",
		middlewareDir:     baseDir + "/infra/middleware/",
		DtoDir:            baseDir + "/adapter/driver/dto/",
		DrivenDir:         baseDir + "/port/driven",
		DriverDir:         baseDir + "/port/driver",
		HandlerDir:        baseDir + "/adapter/driver/handler/",
		RouterDir:         baseDir + "/adapter/driver/router/",
		ManifestDir:       baseDir + "/manifest/",
		I18nDir:           baseDir + "/manifest/i18n/",
		ConfigDir:         baseDir + "/manifest/config/",
		RequestDir:        baseDir + "/adapter/driver/dto/request/",
		ResponseDir:       baseDir + "/adapter/driver/dto/response/",
		UtilsDir:          baseDir + "/infra/utils/",
		TableColumns:      GetTableCol(tableName),
	}
}
