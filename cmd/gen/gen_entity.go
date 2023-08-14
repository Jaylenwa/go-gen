package gen

import (
	"strings"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

// GenEntity 生成entity
func GenEntity(req GenReq) {
	columnStr := "\n"

	for _, v := range req.TableColumns {
		colStr := generateStructFieldForEntity(v)

		columnStr += colStr + "\n"
	}

	str := `package entity
{{TempImports}}
type ` + GetJsonTagFromCase(req.TableName, "Camel") + ` struct {
	` + columnStr +
		`}`

	if strings.Contains(str, "time.Time") {
		str = strings.Replace(str, "{{TempImports}}", "import \"time\"", 1)
	} else {
		str = strings.Replace(str, "{{TempImports}}", "", 1)
	}
	path := req.EntityDir + "/" + req.TableName + "_entity.go"
	if err := gfile.PutContents(path, strings.TrimSpace(str)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

// generateStructFieldForModel 获取字段解析后的 type所索要的类型
func generateStructFieldForEntity(field TableColumn) (colStr string) {
	var fieldName, typeName, jsonTag, node string

	typeName = generateStructFieldTypeNameForEntity(field)

	// 字段名称 如CategoryName
	fieldName = GetJsonTagFromCase(field.Field, "Camel")

	// jsonTag 如 json:"cate_name"
	jsonTag = `json:"` + field.Field + `"`

	// note 如 // 分类名称
	node = " // " + field.Comment

	colStr = fieldName + "    " + typeName + "    " + "`" + " " + jsonTag + "`" + node

	return
}

func generateStructFieldTypeNameForEntity(field TableColumn) string {
	var typeName string
	t, _ := gregex.ReplaceString(`\(.+\)`, "", field.Type)
	t = gstr.Split(gstr.Trim(t), " ")[0]
	t = gstr.ToLower(t)
	switch t {
	case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob":
		typeName = "[]byte"

	case "bit", "int", "int2", "tinyint", "small_int", "smallint", "medium_int", "mediumint", "serial":
		if gstr.ContainsI(field.Type, "unsigned") {
			typeName = "uint"
		} else {
			typeName = "int"
		}

	case "int4", "int8", "big_int", "bigint", "bigserial":
		if gstr.ContainsI(field.Type, "unsigned") {
			typeName = "uint64"
		} else {
			typeName = "int64"
		}

	case "real":
		typeName = "float32"

	case "float", "double", "decimal", "smallmoney", "numeric":
		typeName = "float64"

	case "bool":
		typeName = "bool"

	case "datetime", "timestamp", "date", "time":
		typeName = "time.Time"
	case "json":
		typeName = "string"
	default:
		// Automatically detect its data type.
		switch {
		case strings.Contains(t, "int"):
			typeName = "int"
		case strings.Contains(t, "text") || strings.Contains(t, "char"):
			typeName = "string"
		case strings.Contains(t, "float") || strings.Contains(t, "double"):
			typeName = "float64"
		case strings.Contains(t, "bool"):
			typeName = "bool"
		case strings.Contains(t, "binary") || strings.Contains(t, "blob"):
			typeName = "[]byte"
		case strings.Contains(t, "date") || strings.Contains(t, "time"):
			typeName = "time.Time"
		default:
			typeName = "string"
		}
	}

	return typeName
}
