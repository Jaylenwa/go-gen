package gen

const TempPortDriver = `
package driver

import (
	"context"
	"TempImportPkg/adapter/driver/dto/request"
	"TempImportPkg/adapter/driver/dto/response"
)

//go:generate mockgen -source=./TempSvcNameCaseSnake.go -destination ./mock/TempSvcNameCaseSnake.go -package mock
type TempSvcNameCaseCamelService interface {
	FindTempSvcNameCaseCamelById(ctx context.Context, id int64) (res response.GetTempSvcNameCaseCamelByIdRsp, err error)
	FindTempSvcNameCaseCamelList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []response.GetTempSvcNameCaseCamelListRsp, err error)
	CreateTempSvcNameCaseCamel(ctx context.Context, req request.CreateTempSvcNameCaseCamelReq) (id int64, err error)
	UpdateTempSvcNameCaseCamel(ctx context.Context, id int64, req request.UpdateTempSvcNameCaseCamelReq) (err error)
	DelTempSvcNameCaseCamel(ctx context.Context, id int64) (err error)
}
`
