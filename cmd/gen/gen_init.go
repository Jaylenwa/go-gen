package gen

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

func GenBoot(req GenReq) {
	context := gstr.ReplaceByMap(TempInit, g.MapStrStr{
		"TempImportPkg": viper.Get("server.go_module").(string),
	})

	path := req.InitDir + "/init.go"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
