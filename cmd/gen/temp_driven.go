package gen

const TempPortDriven = `
package driven

import (
	"context"
	"TempImportPkg/infra/po"
	"TempImportPkg/infra/utils/query"
)

//go:generate mockgen -source=./repo_TempSvcNameCaseSnake.go -destination ./mock/repo_TempSvcNameCaseSnake.go -package mock
type TempSvcNameCaseCamelRepo interface {
	FindById(ctx context.Context, id int64) (res *po.TempSvcNameCaseCamel, err error)
	FindByQuery(ctx context.Context, queries []*query.Query) (res *po.TempSvcNameCaseCamel, err error)
	FindList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []*po.TempSvcNameCaseCamel, err error)
	Insert(ctx context.Context, res *po.TempSvcNameCaseCamel) (id int64, err error)
	Update(ctx context.Context, id int64, res *po.TempSvcNameCaseCamel) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

`
