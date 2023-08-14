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

func GenGlobal(req GenReq) {
	context := gstr.ReplaceByMap(TempGlobal, g.MapStrStr{
		"TempImportPkg": viper.Get("server.go_module").(string),
	})

	path := req.GlobalDir + "/global.go"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
