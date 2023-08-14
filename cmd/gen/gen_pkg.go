package gen

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/os/gfile"
	"strings"
)

func GenPkg(req GenReq) {
	genInitConfig(req)
	genUtilsQuery(req)
	genUtilsStruct(req)
	genUtilsValidate(req)
	genErrorHandler(req)
	genRouter(req)
}

func genInitConfig(req GenReq) {
	str := "package config\n\nimport (\n\t\"gopkg.in/yaml.v3\"\n\t\"io/ioutil\"\n\t\"log\"\n\t\"sync\"\n)\n\ntype Project struct {\n\tHost    string `yaml:\"host\"`\n\tPort    int    `yaml:\"port\"`\n\tPortPub int    `yaml:\"port_pub\"`\n\tPortPri int    `yaml:\"port_pri\"`\n}\n\ntype MySQL struct {\n\tUsername        string `yaml:\"username\"`\n\tPassword        string `yaml:\"password\"`\n\tDbHost          string `yaml:\"db_host\"`\n\tDbPort          int    `yaml:\"db_port\"`\n\tDbName          string `yaml:\"db_name\"`\n\tCharset         string `yaml:\"charset\"`\n\tTimeout         string `yaml:\"timeout\"`\n\tTimeoutRead     string `yaml:\"timeout_read\"`\n\tTimeoutWrite    string `yaml:\"timeout_write\"`\n\tMaxOpenConns    int    `yaml:\"max_open_conns\"`\n\tMaxIdleConns    int    `yaml:\"max_idle_conns\"`\n\tConnMaxLifetime int    `yaml:\"conn_max_lifetime\"`\n}\n\ntype Redis struct {\n\tHost       string `yaml:\"host\"`\n\tPort       string `yaml:\"port\"`\n\tPassword   string `yaml:\"password\"`\n\tDB         int    `yaml:\"db\"`\n\tMaxRetries int    `yaml:\"max_retries\"`\n\tPoolSize   int    `yaml:\"pool_size\"`\n}\n\ntype Config struct {\n\tProject Project `yaml:\"project\"`\n\tMySQL   MySQL   `yaml:\"mysql\"`\n\tRedis   Redis   `yaml:\"redis\"`\n}\n\nvar (\n\tconfigOnce sync.Once\n\tconfigImpl *Config\n)\n\nfunc NewConfig() *Config {\n\tconfigOnce.Do(func() {\n\n\t\t// 生产环境\n\t\tconfigFilePath := \"config.yaml\"\n\n\t\tconf := &Config{}\n\t\terr := conf.loadConfig(configFilePath)\n\t\tif err != nil {\n\t\t\tlog.Fatalf(\"load %v failed: %v\", configFilePath, err)\n\t\t\treturn\n\t\t}\n\t})\n\n\treturn configImpl\n}\n\n// loadConfig 加载配置\nfunc (conf *Config) loadConfig(path string) (err error) {\n\n\tfile, err := ioutil.ReadFile(path)\n\tif err != nil {\n\t\tlog.Fatalf(\"load %v failed: %v\", path, err)\n\t}\n\n\terr = yaml.Unmarshal(file, &configImpl)\n\tif err != nil {\n\t\tlog.Fatalf(\"unmarshal yaml file failed: %v\", err)\n\t}\n\n\treturn\n}\n"
	path := req.ConfigDir + "/config.go"
	if err := gfile.PutContents(path, strings.TrimSpace(str)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func genUtilsQuery(req GenReq) {

	str := "package query\n\nimport (\n\t\"fmt\"\n\t\"strconv\"\n\t\"unicode\"\n)\n\n// Query 搜索\ntype Query struct {\n\tKey      string   `json:\"key\"`      // 搜索关键词的键\n\tValue    string   `json:\"value\"`    // 搜索关键词的值\n\tOperator Operator `json:\"operator\"` // 判断条件\n}\n\ntype Operator int32\n\nconst (\n\tOperator_GT    Operator = 0 //大于\n\tOperator_EQUAL Operator = 1 //等于\n\tOperator_LT    Operator = 2 //小于\n\tOperator_NEQ   Operator = 3 //不等于\n\tOperator_LIKE  Operator = 4 //模糊查询\n\tOperator_GTE   Operator = 5 // 大于等于\n\tOperator_LTE   Operator = 6 // 小于等于\n\tOperator_IN    Operator = 7 // in\n)\n\n// Enum value maps for Operator.\nvar (\n\tOperator_name = map[int32]string{\n\t\t0: \"GT\",\n\t\t1: \"EQUAL\",\n\t\t2: \"LT\",\n\t\t3: \"NEQ\",\n\t\t4: \"LIKE\",\n\t\t5: \"GTE\",\n\t\t6: \"LTE\",\n\t\t7: \"IN\",\n\t}\n\tOperator_value = map[string]int32{\n\t\t\"GT\":    0,\n\t\t\"EQUAL\": 1,\n\t\t\"LT\":    2,\n\t\t\"NEQ\":   3,\n\t\t\"LIKE\":  4,\n\t\t\"GTE\":   5,\n\t\t\"LTE\":   6,\n\t\t\"IN\":    7,\n\t}\n)\n\nvar OperatorMap = map[Operator]string{\n\tOperator_GT:    \" > \",\n\tOperator_EQUAL: \" = \",\n\tOperator_LT:    \" < \",\n\tOperator_NEQ:   \" != \",\n\tOperator_LIKE:  \" like \",\n\tOperator_GTE:   \" >= \",\n\tOperator_LTE:   \" <= \",\n\tOperator_IN:    \" in \",\n}\n\n// GenerateQueryCondition 组装 搜索\nfunc GenerateQueryCondition(conditions []*Query) string {\n\tvar condition string\n\tfor k, v := range conditions {\n\t\tif k > 0 {\n\t\t\tcondition += \" and \"\n\t\t}\n\n\t\tif v.Operator == Operator_LIKE {\n\t\t\tcondition += fmt.Sprintf(\"%v%s'%%%v%%'\", v.Key, OperatorMap[v.Operator], CharCheck(v.Value))\n\t\t} else if v.Operator == Operator_IN {\n\t\t\tcondition += fmt.Sprintf(` %s %s (%s)`, v.Key, OperatorMap[v.Operator], v.Value)\n\t\t} else {\n\t\t\t//bool string int\n\t\t\t_, err := strconv.ParseBool(v.Value)\n\t\t\tif err != nil {\n\t\t\t\tcondition += fmt.Sprintf(\"%v%s'%v'\", v.Key, OperatorMap[v.Operator], v.Value)\n\t\t\t} else {\n\t\t\t\tcondition += fmt.Sprintf(\"%v%s%v\", v.Key, OperatorMap[v.Operator], v.Value)\n\t\t\t}\n\t\t}\n\t}\n\n\treturn condition\n}\n\nfunc CharCheck(str string) string {\n\tvar chars []rune\n\tfor _, letter := range str {\n\t\tok, letters := SpecialLetters(letter)\n\t\tif ok {\n\t\t\tchars = append(chars, letters...)\n\t\t} else {\n\t\t\tchars = append(chars, letter)\n\t\t}\n\t}\n\n\treturn string(chars)\n}\n\nfunc SpecialLetters(letter rune) (bool, []rune) {\n\tif unicode.IsPunct(letter) || unicode.IsSymbol(letter) {\n\t\tvar chars []rune\n\t\tchars = append(chars, '\\\\', letter)\n\t\treturn true, chars\n\t}\n\treturn false, nil\n}"

	path := req.QueryDir + "/query.go"
	if err := gfile.PutContents(path, strings.TrimSpace(str)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func genUtilsStruct(req GenReq) {

	str := "package _struct\n\nimport (\n\t\"encoding/json\"\n\t\"errors\"\n\t\"reflect\"\n)\n\n// CopyStruct 将src的值赋值到dst\nfunc CopyStruct(dst interface{}, src interface{}) (err error) {\n\n\tif dst == nil || src == nil {\n\t\terr = errors.New(\"CopyStruct: dst and src must not be nil\")\n\t\treturn\n\t}\n\n\tdstVal := reflect.ValueOf(dst)\n\n\tif dstVal.Kind() != reflect.Ptr {\n\t\terr = errors.New(\"CopyStruct: dst reflect kind need to be a pointer\")\n\t\treturn\n\t}\n\n\tif dstVal.IsNil() {\n\t\terr = errors.New(\"CopyStruct: dstVal.IsNil()\")\n\t\treturn\n\t}\n\n\tbys, err := json.Marshal(src)\n\tif err != nil {\n\t\treturn\n\t}\n\n\terr = json.Unmarshal(bys, dst)\n\n\treturn\n}"
	path := req.StructDir + "/struct.go"
	if err := gfile.PutContents(path, strings.TrimSpace(str)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func genUtilsValidate(req GenReq) {
	str := "package validate\n\nimport (\n\t\"reflect\"\n\t\"regexp\"\n\n\t\"github.com/go-playground/validator/v10\"\n\t\"github.com/pkg/errors\"\n)\n\nvar validate *validator.Validate\n\nvar (\n\t// customDataTag is default data tag name\n\tcustomDataTag = \"json\"\n\t// customErrTag is default custom tag name\n\tcustomErrTag = \"err_info\"\n)\n\nfunc init() {\n\tvalidate = validator.New()\n\t// 注册自定义错误\n\tif err := validate.RegisterValidation(\"checkSpecialChar\", checkSpecialChar); err != nil {\n\t\tpanic(err)\n\t}\n}\n\n// SetCustomDataTag set custom data tag name\nfunc SetCustomDataTag(tag string) {\n\tcustomDataTag = tag\n}\n\n// SetCustomErrTag set custom err tag name\nfunc SetCustomErrTag(tag string) {\n\tcustomErrTag = tag\n}\n\n// Validate is validate a struct exposed fields\nfunc Validate(val interface{}) error {\n\terr := validate.Struct(val)\n\tif err == nil {\n\t\treturn nil\n\t}\n\n\tfor _, err := range err.(validator.ValidationErrors) {\n\t\treturn wrapErr(val, err)\n\t}\n\n\treturn nil\n}\n\n// wrapErr is wrap err\nfunc wrapErr(val interface{}, err validator.FieldError) error {\n\tt := reflect.TypeOf(val)\n\tif t.Kind() == reflect.Ptr {\n\t\tt = t.Elem()\n\t}\n\n\tf, ok := t.FieldByName(err.Field())\n\tif !ok {\n\t\treturn errors.Errorf(\"param %s must %s %s\", err.Field(), err.Tag(), err.Param())\n\t}\n\n\terrTag := f.Tag.Get(customErrTag)\n\tif errTag == \"\" {\n\t\treturn errors.Errorf(\"param %s must %s %s\", f.Tag.Get(customDataTag), err.Tag(), err.Param())\n\t}\n\n\treturn errors.Errorf(\"%s:%s\", f.Tag.Get(customDataTag), errTag)\n}\n\n// checkSpecialChar 校验特殊字符\nfunc checkSpecialChar(f validator.FieldLevel) bool {\n\tvalue := f.Field().String()\n\tif value == \"\" {\n\t\treturn true\n\t}\n\n\tflag, err := regexp.MatchString(\"^([A-Za-z0-9@_.-]{1,128})$\", value)\n\n\tif err != nil {\n\t\treturn false\n\t}\n\n\treturn flag\n}"
	path := req.ValidateDir + "/validate.go"
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
	str := "package driver\n\nimport \"github.com/gin-gonic/gin\"\n\n// HttpRouterInterface 路由公共接口\ntype HttpRouterInterface interface {\n\t// RegisterRouterPublic 注册外部API\n\tRegisterRouterPublic(engine *gin.RouterGroup)\n\n\t// RegisterRouterPrivate 注册内部API\n\tRegisterRouterPrivate(engine *gin.RouterGroup)\n}\n"
	path := req.RouterDir + "/router.go"
	if err := gfile.PutContents(path, strings.TrimSpace(str)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}
