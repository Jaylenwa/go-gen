package gen

import (
	"strings"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/os/gfile"
)

func GenPkg(req GenReq) {
	genInitConfig(req)
	genErrorHandler(req)
	genRouter(req)
	genLog(req)
}

func genLog(req GenReq) {
	str := `
	package log

	import "github.com/sirupsen/logrus"

	func InitLog() *logrus.Logger {
		return logrus.New()
	}
	`
	path := req.LogDir + "/log.go"
	if err := gfile.PutContents(path, strings.TrimSpace(str)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func genInitConfig(req GenReq) {

	str := "package config\n\nimport (\n\t\"gopkg.in/yaml.v3\"\n\t\"os\"\n\t\"log\"\n\t\"sync\"\n)\n\ntype Project struct {\n\tHost    string `yaml:\"host\"`\n\tPort    int    `yaml:\"port\"`\n}\n\ntype MySQL struct {\n\tUsername        string `yaml:\"username\"`\n\tPassword        string `yaml:\"password\"`\n\tDbHost          string `yaml:\"db_host\"`\n\tDbPort          int    `yaml:\"db_port\"`\n\tDbName          string `yaml:\"db_name\"`\n\tCharset         string `yaml:\"charset\"`\n\tTimeout         string `yaml:\"timeout\"`\n\tTimeoutRead     string `yaml:\"timeout_read\"`\n\tTimeoutWrite    string `yaml:\"timeout_write\"`\n\tMaxOpenConns    int    `yaml:\"max_open_conns\"`\n\tMaxIdleConns    int    `yaml:\"max_idle_conns\"`\n\tConnMaxLifetime int    `yaml:\"conn_max_lifetime\"`\n\tLogLevel int    `yaml:\"log_level\"`\n}\n\ntype Redis struct {\n\tHost       string `yaml:\"host\"`\n\tPort       string `yaml:\"port\"`\n\tPassword   string `yaml:\"password\"`\n\tDB         int    `yaml:\"db\"`\n\tMaxRetries int    `yaml:\"max_retries\"`\n\tPoolSize   int    `yaml:\"pool_size\"`\n}\n\ntype Config struct {\n\tProject Project `yaml:\"project\"`\n\tMySQL   MySQL   `yaml:\"mysql\"`\n\tRedis   Redis   `yaml:\"redis\"`\n}\n\nvar (\n\tconfigOnce sync.Once\n\tconfigImpl *Config\n)\n\nfunc InitConfig() *Config {\n\tconfigOnce.Do(func() {\n\n\t\t// 生产环境\n\t\tconfigFilePath := \"manifest/config/config.yaml\"\n\n\t\tconf := &Config{}\n\t\terr := conf.loadConfig(configFilePath)\n\t\tif err != nil {\n\t\t\tlog.Fatalf(\"load %v failed: %v\", configFilePath, err)\n\t\t\treturn\n\t\t}\n\t})\n\n\treturn configImpl\n}\n\n// loadConfig 加载配置\nfunc (conf *Config) loadConfig(path string) (err error) {\n\n\tfile, err := os.ReadFile(path)\n\tif err != nil {\n\t\tlog.Fatalf(\"load %v failed: %v\", path, err)\n\t}\n\n\terr = yaml.Unmarshal(file, &configImpl)\n\tif err != nil {\n\t\tlog.Fatalf(\"unmarshal yaml file failed: %v\", err)\n\t}\n\n\treturn\n}\n"
	path := req.InitConfigDir + "/config.go"
	if err := gfile.PutContents(path, strings.TrimSpace(str)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func genErrorHandler(req GenReq) {
	str := "package middleware\n\nimport (\n\t\"net/http\"\n\n\t\"github.com/gin-gonic/gin\"\n)\n\nfunc ErrorHandlerMiddleware() gin.HandlerFunc {\n\treturn func(c *gin.Context) {\n\t\t// 错误处理\n\t\tdefer func() {\n\t\t\tfor _, err := range c.Errors {\n\t\t\t\tc.AbortWithStatusJSON(c.Writer.Status(), gin.H{\n\t\t\t\t\t\"code\":    c.Writer.Status(),\n\t\t\t\t\t\"message\": http.StatusText(c.Writer.Status()),\n\t\t\t\t\t\"cause\":   err.Error(),\n\t\t\t\t})\n\t\t\t\treturn\n\t\t\t}\n\t\t}()\n\t\tc.Writer.Header().Set(\"Content-Type\", \"application/json; charset=utf-8\")\n\t\tc.Next()\n\t}\n}"
	path := req.middlewareDir + "/error_handler.go"
	if err := gfile.PutContents(path, strings.TrimSpace(str)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func genRouter(req GenReq) {
	str := `
	package router

	import "github.com/gin-gonic/gin"

	type CommonRouter interface {
		InitRouter(Router *gin.RouterGroup)
	}
	`
	path := req.RouterDir + "/router.go"
	if err := gfile.PutContents(path, strings.TrimSpace(str)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
