package gen

import (
	"strings"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

func GenREADME(req GenReq) {

	context := gstr.ReplaceByMap(TempREADME, g.MapStrStr{})

	path := req.BaseDir + "/README.MD"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
