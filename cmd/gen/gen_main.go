package gen

import (
	"github.com/spf13/viper"
	"strings"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

func GenMain(req GenReq) {
	context := gstr.ReplaceByMap(TempMain, g.MapStrStr{
		"TempImportPkg":         viper.Get("server.go_module").(string),
		"TempSvcNameCaseCamel":  GetJsonTagFromCase(req.TableName, "CamelLower"), // 表名 大驼峰
		"TempSvcNameCamelLower": GetJsonTagFromCase(req.TableName, "Camel"),
	})

	path := req.BaseDir + "/main.go"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
