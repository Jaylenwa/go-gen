package gen

const TempResponse = `
package response
{{TempImports}}

// GetTempSvcNameCaseCamelByIdRsp
type GetTempSvcNameCaseCamelByIdRsp struct {
	{{findRsp}}
}

// GetTempSvcNameCaseCamelByQueryRsp
type GetTempSvcNameCaseCamelByQueryRsp struct {
	{{findRsp}}
}

// GetTempSvcNameCaseCamelListRsp
type GetTempSvcNameCaseCamelListRsp struct {
	{{findRsp}}
}
`
