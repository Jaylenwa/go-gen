package gen

import (
	"gencode/util"
	"github.com/gogf/gf-cli/v2/library/mlog"
	"os"
)

func Run(host, user, password, port, db, table, serverName string) {
	// 1. 获取表完整结构信息
	InitDB(host, port, user, password, db)

	genReq := GenInit(serverName, table)

	mlog.Print("auto gen code start...")

	CreateDir(genReq)

	GenMod(genReq)

	GenPkg(genReq)

	GenEntity(genReq)

	GenModel(genReq)

	GenMain(genReq)

	GenMysql(genReq)

	GenGlobal(genReq)

	GenDto(genReq)

	GenPortDriver(genReq)

	GenPortDriven(genReq)

	GenHandler(genReq)

	GenBoot(genReq)

	GenService(genReq)

	GenRepo(genReq)

	GenConfig(genReq)

	// 自动化文档
	//GenApiDoc(genReq)

	// 10.格式化代码
	util.GoFmt(genReq.BaseDir)

	mlog.Print("done!")
}

// CreateDir 创建需要的文件夹
func CreateDir(req GenReq) {
	_ = os.MkdirAll(req.BootDir, os.ModePerm)
	_ = os.MkdirAll(req.EntityDir, os.ModePerm)
	_ = os.MkdirAll(req.ServiceDir, os.ModePerm)
	_ = os.MkdirAll(req.GlobalDir, os.ModePerm)
	_ = os.MkdirAll(req.ConfigDir, os.ModePerm)
	_ = os.MkdirAll(req.ConstDir, os.ModePerm)
	_ = os.MkdirAll(req.MysqlDir, os.ModePerm)
	_ = os.MkdirAll(req.RepositoryPoDir, os.ModePerm)
	_ = os.MkdirAll(req.RepositoryImplDir, os.ModePerm)
	_ = os.MkdirAll(req.QueryDir, os.ModePerm)
	_ = os.MkdirAll(req.DtoDir, os.ModePerm)
	_ = os.MkdirAll(req.middlewareDir, os.ModePerm)
	_ = os.MkdirAll(req.DriverDir, os.ModePerm)
	_ = os.MkdirAll(req.DrivenDin, os.ModePerm)
	_ = os.MkdirAll(req.ConcurrentMap, os.ModePerm)

}
