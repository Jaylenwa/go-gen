package cmd

import (
	"gencode/cmd/gen"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	addr, user, pwd, port, db, table, serverName string
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "automatically generate go files for ORM model, service, repository, handler",
	Long:  `automatically generate go files for ORM model, service, repository, handler`,
	Run: func(cmd *cobra.Command, args []string) {
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
	genCmd.Flags().StringVarP(&serverName, "serverName", "s", "", "Enter project server name")
}

func help() {
	mlog.Print(gstr.TrimLeft(`
USAGE
    gencode gen [OPTION]

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
    gencode gen -a 127.0.0.1 -u root -pwd root -p 3306 -d dbName -t tableName -s microName
`))
}
