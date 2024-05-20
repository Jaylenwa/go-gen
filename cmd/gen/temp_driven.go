package gen

const TempPortDriven = `
package driven

import (
	"context"
	"TempImportPkg/infra/po"
)

//go:generate mockgen -source=./TempSvcNameCaseSnake.go -destination ./mock/TempSvcNameCaseSnake.go -package mock
type TempSvcNameCaseCamelRepo interface {
	FindById(ctx context.Context, id int64) (res po.TempSvcNameCaseCamel, err error)
	FindByQuery(ctx context.Context, filter map[string]interface{}, args ...interface{}) (res po.TempSvcNameCaseCamel, err error)
	FindList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []po.TempSvcNameCaseCamel, err error)
	Count(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, err error)
	Insert(ctx context.Context, res po.TempSvcNameCaseCamel) (id int64, err error)
	Update(ctx context.Context, id int64, res po.TempSvcNameCaseCamel) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
`
