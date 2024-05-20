package gen

const TempRequest = `
package request
{{TempImports}}

// GetTempSvcNameCaseCamelListReq
type GetTempSvcNameCaseCamelListReq struct {
	{{findListDto}}
}

// CreateTempSvcNameCaseCamelReq
type CreateTempSvcNameCaseCamelReq struct {
	{{createDto}}
}

// DelTempSvcNameCaseCamelReq
type DelTempSvcNameCaseCamelReq struct {
	{{delDto}}
}

// UpdateTempSvcNameCaseCamelReq
type UpdateTempSvcNameCaseCamelReq struct {
	{{updateDto}}
}

// FindTempSvcNameCaseCamelByIdReq
type FindTempSvcNameCaseCamelByIdReq struct {
	{{findByIdDto}}
}
`
