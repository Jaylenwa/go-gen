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
	SrvName           string
	TableKey          string
	BaseDir           string
	BootDir           string
	EntityDir         string
	ServiceDir        string
	GlobalDir         string
	InfraDir          string
	ConfigDir         string
	ConstDir          string
	MysqlDir          string
	RepositoryPoDir   string
	RepositoryImplDir string
	QueryDir          string
	HttpClientDir     string
	//StructDir         string
	ValidateDir   string
	DtoDir        string
	middlewareDir string
	DrivenDin     string
	DriverDir     string
	HandlerDir    string
	RouterDir     string

	TableColumns []TableColumn
}

func GenInit(srvName, tableName string) GenReq {
	codePath := viper.Get("server.code_path").(string)
	baseDir := codePath + "/" + viper.Get("server.go_module").(string)

	return GenReq{
		BaseDir:           baseDir,
		TableName:         tableName,
		TableKey:          viper.Get("mysql.key").(string),
		SrvName:           srvName,
		BootDir:           baseDir + "/boot/",
		EntityDir:         baseDir + "/domain/entity/",
		ServiceDir:        baseDir + "/domain/service/",
		GlobalDir:         baseDir + "/global/",
		InfraDir:          baseDir + "/infra",
		ConfigDir:         baseDir + "/infra/config/",
		ConstDir:          baseDir + "/infra/const/",
		MysqlDir:          baseDir + "/infra/db/mysql/",
		RepositoryPoDir:   baseDir + "/infra/po/",
		RepositoryImplDir: baseDir + "/adapter/driven/",
		QueryDir:          baseDir + "/infra/utils/query/",
		HttpClientDir:     baseDir + "/infra/utils/httpclient/",
		//StructDir:         baseDir + "/infra/utils/struct/",
		ValidateDir:   baseDir + "/infra/utils/validate/",
		middlewareDir: baseDir + "/infra/middleware/",
		DtoDir:        baseDir + "/adapter/driver/dto/",
		DrivenDin:     baseDir + "/port/driven",
		DriverDir:     baseDir + "/port/driver",
		HandlerDir:    baseDir + "/adapter/driver/",
		RouterDir:     baseDir + "/adapter/driver/",
		TableColumns:  GetTableCol(tableName),
	}
}
