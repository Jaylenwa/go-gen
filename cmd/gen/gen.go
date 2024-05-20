package gen

import (
	"gogen/util"
	"os"

	"github.com/gogf/gf-cli/v2/library/mlog"
)

func Run(host, user, password, port, db, table string) {
	// 1. 获取表完整结构信息
	InitDB(host, port, user, password, db)

	genReq := GenInit(table)

	mlog.Print("auto gen code start...")

	CreateDir(genReq)

	GenMod(genReq)

	GenPkg(genReq)

	GenEntity(genReq)

	GenModel(genReq)

	GenMain(genReq)

	GenMysql(genReq)

	GenGlobal(genReq)

	GenRequest(genReq)

	GenResponse(genReq)

	GenPortDriver(genReq)

	GenPortDriven(genReq)

	GenHandler(genReq)

	GenBoot(genReq)

	GenService(genReq)

	GenRepo(genReq)

	GenREADME(genReq)

	GenConfig(genReq)

	// 自动化文档
	//GenApiDoc(genReq)

	// 10.格式化代码
	util.GoFmt(genReq.BaseDir)

	mlog.Print("done!")
}

// CreateDir 创建需要的文件夹
func CreateDir(req GenReq) {
	_ = os.MkdirAll(req.InitDir, os.ModePerm)
	_ = os.MkdirAll(req.EntityDir, os.ModePerm)
	_ = os.MkdirAll(req.ServiceDir, os.ModePerm)
	_ = os.MkdirAll(req.GlobalDir, os.ModePerm)
	_ = os.MkdirAll(req.ConfigDir, os.ModePerm)
	_ = os.MkdirAll(req.EnumsDir, os.ModePerm)
	_ = os.MkdirAll(req.DBDir, os.ModePerm)
	_ = os.MkdirAll(req.LogDir, os.ModePerm)
	_ = os.MkdirAll(req.RepositoryPoDir, os.ModePerm)
	_ = os.MkdirAll(req.RepositoryImplDir, os.ModePerm)
	_ = os.MkdirAll(req.DtoDir, os.ModePerm)
	_ = os.MkdirAll(req.middlewareDir, os.ModePerm)
	_ = os.MkdirAll(req.DriverDir, os.ModePerm)
	_ = os.MkdirAll(req.DrivenDir, os.ModePerm)
	_ = os.MkdirAll(req.ManifestDir, os.ModePerm)
	_ = os.MkdirAll(req.InitConfigDir, os.ModePerm)
	_ = os.MkdirAll(req.I18nDir, os.ModePerm)
	_ = os.MkdirAll(req.RequestDir, os.ModePerm)
	_ = os.MkdirAll(req.ResponseDir, os.ModePerm)
	_ = os.MkdirAll(req.UtilsDir, os.ModePerm)
}
