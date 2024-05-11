package cmd

import (
	"gogen/cmd/gen"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	addr, user, pwd, port, db, table, serverName, goModule, goVersion, codePath string
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "automatically generate go files for ORM model, service, repository, handler",
	Long:  `automatically generate go files for ORM model, service, repository, handler`,
	Run: func(cmd *cobra.Command, args []string) {
		println(1)
		i := viper.Get("server.go_version")
		println(i.(string))
		println(2)
		addr = viper.Get("mysql.addr").(string)
		user = viper.Get("mysql.user").(string)
		pwd = viper.Get("mysql.pwd").(string)
		port = viper.Get("mysql.port").(string)
		db = viper.Get("mysql.db").(string)
		table = viper.Get("mysql.table").(string)
		serverName = viper.Get("server.server_name").(string)
		if len(addr) == 0 || len(user) == 0 || len(port) == 0 || len(db) == 0 || len(table) == 0 || len(serverName) == 0 {
			help()
			return
		}
		gen.Run(addr, user, pwd, port, db, table, serverName)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1", "Enter MySQL addr")
	genCmd.Flags().StringVarP(&user, "user", "u", "root", "Enter MySQL user")
	genCmd.Flags().StringVarP(&pwd, "pwd", "", "root", "Enter MySQL password")
	genCmd.Flags().StringVarP(&port, "port", "p", "3306", "Enter MySQL port")
	genCmd.Flags().StringVarP(&db, "db", "d", "", "Enter MySQL database")
	genCmd.Flags().StringVarP(&table, "table", "t", "", "Enter MySQL table")
	genCmd.Flags().StringVarP(&goVersion, "key", "", "id", "Enter table primary key")
	genCmd.Flags().StringVarP(&serverName, "server_name", "s", "", "Enter project server name")
	genCmd.Flags().StringVarP(&goModule, "go_module", "", "tmp", "Enter project go module name")
	genCmd.Flags().StringVarP(&codePath, "code_path", "", "./tmp", "Enter project code generation path")
	genCmd.Flags().StringVarP(&goVersion, "go_version", "", "1.20", "Enter project go version")
	if len(viper.AllKeys()) == 0 {
		viper.BindPFlag("mysql.addr", genCmd.Flags().Lookup("addr"))
		viper.BindPFlag("mysql.user", genCmd.Flags().Lookup("user"))
		viper.BindPFlag("mysql.pwd", genCmd.Flags().Lookup("pwd"))
		viper.BindPFlag("mysql.port", genCmd.Flags().Lookup("port"))
		viper.BindPFlag("mysql.db", genCmd.Flags().Lookup("db"))
		viper.BindPFlag("mysql.table", genCmd.Flags().Lookup("table"))
		viper.BindPFlag("mysql.key", genCmd.Flags().Lookup("key"))
		viper.BindPFlag("server.server_name", genCmd.Flags().Lookup("server_name"))
		viper.BindPFlag("server.go_module", genCmd.Flags().Lookup("go_module"))
		viper.BindPFlag("server.code_path", genCmd.Flags().Lookup("code_path"))
		viper.BindPFlag("server.go_version", genCmd.Flags().Lookup("go_version"))
	}
}

func help() {
	mlog.Print(gstr.TrimLeft(`
USAGE
    gogen gen [OPTION]

ARGUMENT
    OPTION
	-a	Enter MySQL addr
	-u	Enter MySQL user
	-pwd	Enter MySQL password
	-p	Enter MySQL port
	-d	Enter MySQL database
	-t	Enter MySQL table
	-s	Enter project server name

EXAMPLES
    gogen gen -a 127.0.0.1 -u root -pwd root -p 3306 -d dbName -t tableName -s serverName
`))
}
