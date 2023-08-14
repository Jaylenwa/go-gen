package gen

const TempDto = `
package dto
{{TempImports}}
// 请求对象
type (
	// GetTempSvcNameCaseCamelListReq 查询TempSvcNameCaseCamel列表 请求对象
	GetTempSvcNameCaseCamelListReq struct {
		{{findListDto}}
	}

	// CreateTempSvcNameCaseCamelReq 创建TempSvcNameCaseCamel 请求对象
	CreateTempSvcNameCaseCamelReq struct {
		{{createDto}}
	}

	// DelTempSvcNameCaseCamelReq 删除 请求对象
	DelTempSvcNameCaseCamelReq struct {
		{{delDto}}
	}

	// UpdateTempSvcNameCaseCamelReq 修改TempSvcNameCaseCamel 请求对象
	UpdateTempSvcNameCaseCamelReq struct {
		{{updateDto}}
	}

	// FindTempSvcNameCaseCamelByIdReq 查询 请求对象
	FindTempSvcNameCaseCamelByIdReq struct {
		{{findByIdDto}}
	}
)

// 输出对象
type (
	// GetTempSvcNameCaseCamelByIdRsp 查询TempSvcNameCaseCamelById 返回对象
	GetTempSvcNameCaseCamelByIdRsp struct {
		{{findRsp}}
	}

	// GetTempSvcNameCaseCamelByQueryRsp 查询TempSvcNameCaseCamelByQuery 返回对象
	GetTempSvcNameCaseCamelByQueryRsp struct {
		{{findRsp}}
	}

	// GetTempSvcNameCaseCamelListRsp 查询列表 返回对象
	GetTempSvcNameCaseCamelListRsp struct {
		{{findRsp}}
	}
)

`
