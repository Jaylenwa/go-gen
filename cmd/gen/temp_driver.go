package gen

const TempPortDriver = `

package driver

import (
	"context"
	"TempImportPkg/adapter/driver/dto"
)

// go:generate mockgen -source=./svc_TempSvcNameCaseSnake.go -destination ./mock/svc_TempSvcNameCaseSnake.go -package mock
type TempSvcNameCaseCamelService interface {
	FindTempSvcNameCaseCamelById(ctx context.Context, id int64) (res dto.GetTempSvcNameCaseCamelByIdRsp, err error)
	FindTempSvcNameCaseCamelList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []dto.GetTempSvcNameCaseCamelListRsp, err error)
	CreateTempSvcNameCaseCamel(ctx context.Context, req dto.CreateTempSvcNameCaseCamelReq) (id int64, err error)
	UpdateTempSvcNameCaseCamel(ctx context.Context, id int64, req dto.UpdateTempSvcNameCaseCamelReq) (err error)
	DelTempSvcNameCaseCamel(ctx context.Context, id int64) (err error)
}

`
