package gen

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf-cli/v2/library/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
)

func GenDto(req GenReq) {
	path := req.DtoDir + "/" + req.TableName + "_dto.go"

	context := gstr.ReplaceByMap(TempDto, g.MapStrStr{
		"{{findListDto}}":   genList(req),
		"{{createDto}}":     genAdd(req),
		"{{delDto}}":        genDel(req),
		"{{updateDto}}":     genUpdate(req),
		"{{findByIdDto}}":   genFindById(req),
		"{{updateByIdDto}}": genFindById(req),
		"{{findQueryDto}}":  genFindQuery(req),
		"{{findAllDto}}":    genFindQuery(req),
		"{{createRsp}}":     genCreateRsp(req),
		"{{findRsp}}":       genFindRsp(req),
	})

	context2 := gstr.ReplaceByMap(context, g.MapStrStr{
		"TempSvcNameCaseCamel": GetJsonTagFromCase(req.TableName, "Camel"),
	})

	if strings.Contains(context2, "time.Time") {
		context2 = strings.Replace(context2, "{{TempImports}}", "import \"time\"", 1)
	} else {
		context2 = strings.Replace(context2, "{{TempImports}}", "", 1)
	}

	if err := gfile.PutContents(path, strings.TrimSpace(context2)); err != nil {
		mlog.Fatalf("writing content to '%s' failed: %v", path, err)
	} else {
		utils.GoFmt(path)
		mlog.Print("generated:", path)
	}
}

func genAdd(req GenReq) (columnStr string) {
	columnStr = ""
	for _, v := range req.TableColumns {
		if v.Field == "created_at" || v.Field == "created_by" || v.Field == "updated_at" || v.Field == "updated_by" || v.Field == "deleted_at" || v.Field == "deleted_by" || v.Field == req.TableKey {
			continue
		}
		colStr := generateStructFieldForDto(v) // 单行字段
		columnStr += colStr + "\n"
	}

	return columnStr
}

func genList(req GenReq) (columnStr string) {
	columnStr = "\t\tLimit  int `validate:\"gte=0\" form:\"limit\" json:\"limit\"`\n\t\tOffset int `form:\"offset\" json:\"offset\"`"

	return columnStr
}

func genDel(req GenReq) (columnStr string) {
	columnStr = ""
	for _, v := range req.TableColumns {
		if v.Field == req.TableKey {
			colStr := generateStructFieldForDtoReq(v) // 单行字段
			columnStr += colStr + "\n"
			break
		}
	}

	return columnStr
}

func genUpdate(req GenReq) (columnStr string) {
	columnStr = ""
	for _, v := range req.TableColumns {
		if v.Field == "created_at" || v.Field == "created_by" || v.Field == "updated_at" || v.Field == "updated_by" || v.Field == "deleted_at" || v.Field == "deleted_by" {
			continue
		}
		colStr := generateStructFieldForDtoReq(v) // 单行字段
		columnStr += colStr + "\n"
	}

	return columnStr
}

func genFindById(req GenReq) (columnStr string) {
	columnStr = ""
	for _, v := range req.TableColumns {
		if v.Field == req.TableKey {
			colStr := generateStructFieldForDtoReq(v) // 单行字段
			columnStr += colStr + "\n"
			break
		}
	}

	return columnStr
}

func genFindQuery(req GenReq) (columnStr string) {
	return "Query []*types.Query `json:\"query\"`"
}

func genCreateRsp(req GenReq) (columnStr string) {
	columnStr = ""
	for _, v := range req.TableColumns {
		if v.Field == req.TableKey {
			colStr := generateStructFieldForDto(v) // 单行字段
			columnStr += colStr + "\n"
			break
		}
	}

	return columnStr
}

func genFindRsp(req GenReq) (columnStr string) {
	columnStr = ""
	for _, v := range req.TableColumns {
		if v.Field == "deleted_at" || v.Field == "deleted_by" {
			continue
		}
		colStr := generateStructFieldForDto(v) // 单行字段
		columnStr += colStr + "\n"
	}

	return columnStr
}

// generateStructFieldForModel 获取字段解析后的 type所索要的类型
func generateStructFieldForDto(field TableColumn) (colStr string) {
	var fieldName, typeName, jsonTag, node string

	typeName = generateStructFieldTypeNameForEntity(field)

	// 字段名称 如CategoryName
	fieldName = GetJsonTagFromCase(field.Field, "Camel")

	// jsonTag 如 json:"cate_name"
	jsonTag = `json:"` + field.Field + `"`

	// note 如 // 分类名称
	if field.Comment != "" {
		node = " // " + field.Comment
	}

	colStr = fieldName + "    " + typeName + "    " + "`" + " " + jsonTag + "`" + node

	return
}

// generateStructFieldForDtoReq 获取字段解析后的 type所索要的类型
func generateStructFieldForDtoReq(field TableColumn) (colStr string) {
	var fieldName, typeName, jsonTag, node string

	typeName = generateStructFieldTypeNameForEntity(field)

	// 字段名称 如CategoryName
	fieldName = GetJsonTagFromCase(field.Field, "Camel")

	// jsonTag 如 json:"cate_name"
	jsonTag = `json:"` + field.Field + `"`
	if gstr.ContainsI(field.Key, "pri") {
		jsonTag = `validate:"required" uri:"` + field.Field + `" json:"` + field.Field + `"`
	}

	// note 如 // 分类名称
	if field.Comment != "" {
		node = " // " + field.Comment
	}

	colStr = fieldName + "    " + typeName + "    " + "`" + " " + jsonTag + "`" + node

	return
}
