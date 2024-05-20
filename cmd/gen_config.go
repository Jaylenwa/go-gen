package cmd

import (
	"strings"

	"github.com/gogf/gf-cli/library/mlog"
	"github.com/gogf/gf-cli/library/utils"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	"github.com/spf13/cobra"
)

var genConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "atomatically generate configuration files",
	Long:  `automatically generate configuration files`,
	Run: func(cmd *cobra.Command, args []string) {
		genenerateConfig()
	},
}

func init() {
	rootCmd.AddCommand(genConfigCmd)
}

// 在当前目录下生成配置文件
func genenerateConfig() {
	context := gstr.ReplaceByMap(TempConfig, g.MapStrStr{})
	path := ".gen-cli.yaml"
	if err := gfile.PutContents(path, strings.TrimSpace(context)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

const TempConfig = `
server:
  # go.mod 模块名称【必填】
  go_module: "test"
  # 自动化代码输出目录
  code_path: "./tmp"
  # go 版本
  go_version: "1.20"
mysql:
  # 数据库连接地址【必填】
  addr: "127.0.0.1"
  # 数据库端口【必填】
  port: "3306"
  # 数据库用户【必填】
  user: "root"
  # 数据库用户密码【必填】
  pwd: "pwd"
  # 数据库名称【必填】
  db: "test"
  # 数据库表名【必填】
  table: "test"
  # 主键
  key: "id"
`
