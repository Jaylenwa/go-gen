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

func GenConfig(req GenReq) {
	context := gstr.ReplaceByMap(TempConfig, g.MapStrStr{
		"TempDBAddr": viper.Get("mysql.addr").(string),
		"TempDBPort": viper.Get("mysql.port").(string),
		"TempDBUser": viper.Get("mysql.user").(string),
		"TempDBPwd":  viper.Get("mysql.pwd").(string),
		"TempDBName": viper.Get("mysql.db").(string),
	})

	path := req.BaseDir + "/config.yaml"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
