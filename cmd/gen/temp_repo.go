package gen

const TempRepo = `
package driven

import (
	"context"
	"TempImportPkg/infra/po"

	portDriven "TempImportPkg/port/driven"
)

type TempSvcNameCamelLowerRepo struct {}

var _ portDriven.TempSvcNameCaseCamelRepo = &TempSvcNameCamelLowerRepo{}

func NewTempSvcNameCaseCamelRepo() portDriven.TempSvcNameCaseCamelRepo {
	return &TempSvcNameCamelLowerRepo{}
}

func (repo *TempSvcNameCamelLowerRepo) FindById(ctx context.Context, id int64) (TempSvcNameCamelLower po.TempSvcNameCaseCamel, err error) {
	err = getDb().Where("id = ?", id).First(&TempSvcNameCamelLower).Error
	return
}

func (repo *TempSvcNameCamelLowerRepo) FindByQuery(ctx context.Context, filter map[string]interface{}, args ...interface{}) (res po.TempSvcNameCaseCamel, err error) {
	dbQuery := getDb().Model(&po.TempSvcNameCaseCamel{}).Where(filter)
	if len(args) >= 2 {
		dbQuery = dbQuery.Where(args[0], args[1:]...)
	} else if len(args) >= 1 {
		dbQuery = dbQuery.Where(args[0])
	}

	err = dbQuery.First(&res).Error
	return
}

func (repo *TempSvcNameCamelLowerRepo) FindList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []po.TempSvcNameCaseCamel, err error) {
	limit := 10
	offset := 0

	if v, ok := filter["limit"]; ok {
		limit = int(v.(float64))
		delete(filter, "limit")
	}

	if v, ok := filter["offset"]; ok {
		offset = int(v.(float64))
		delete(filter, "offset")
	}

	dbQuery := getDb().Model(&po.TempSvcNameCaseCamel{}).Where(filter)

	if len(args) >= 2 {
		dbQuery = dbQuery.Where(args[0], args[1:]...)
	} else if len(args) >= 1 {
		dbQuery = dbQuery.Where(args[0])
	}

	dbQuery = dbQuery.Count(&total)

	err = dbQuery.Limit(limit).Offset(offset).Find(&res).Error
	return
}

func (repo *TempSvcNameCamelLowerRepo) Count(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, err error) {
	dbQuery := getDb().Model(&po.TempSvcNameCaseCamel{}).Where(filter)
	if len(args) >= 2 {
		dbQuery = dbQuery.Where(args[0], args[1:]...)
	} else if len(args) >= 1 {
		dbQuery = dbQuery.Where(args[0])
	}

	err = dbQuery.Count(&total).Error
	return
}

func (repo *TempSvcNameCamelLowerRepo) Insert(ctx context.Context, TempSvcNameCamelLower po.TempSvcNameCaseCamel) (id int64, err error) {
	err = getTx(ctx).Create(&TempSvcNameCamelLower).Error
	if err != nil {
		return
	}

	id = TempSvcNameCamelLower.Id
	return
}

func (repo *TempSvcNameCamelLowerRepo) Update(ctx context.Context, id int64, TempSvcNameCamelLower po.TempSvcNameCaseCamel) (err error) {
	err = getTx(ctx).Model(&po.TempSvcNameCaseCamel{}).Where("id = ?", id).Updates(&TempSvcNameCamelLower).Error
	return
}

func (repo *TempSvcNameCamelLowerRepo) Delete(ctx context.Context, id int64) (err error) {
	err = getTx(ctx).Where("id = ?", id).Delete(&po.TempSvcNameCaseCamel{}).Error
	return
}
`
