package gen

const TempRepo = `
package adapterDriven

import (
	"context"
	"TempImportPkg/global"
	"TempImportPkg/infra/po"
	portDriven "TempImportPkg/port/driven"
	"gorm.io/gorm"
	"sync"
)

var (
	TempSvcNameCamelLowerRepoOnce sync.Once
	TempSvcNameCamelLowerRepoImpl portDriven.TempSvcNameCaseCamelRepo
)

type TempSvcNameCamelLowerRepo struct {
	db *gorm.DB
}

var _ portDriven.TempSvcNameCaseCamelRepo = &TempSvcNameCamelLowerRepo{}

func NewTempSvcNameCaseCamelRepo() portDriven.TempSvcNameCaseCamelRepo {
	TempSvcNameCamelLowerRepoOnce.Do(func() {
		TempSvcNameCamelLowerRepoImpl = &TempSvcNameCamelLowerRepo{
			db: global.GDB,
		}
	})
	return TempSvcNameCamelLowerRepoImpl
}

func (repo *TempSvcNameCamelLowerRepo) FindById(ctx context.Context, id int64) (TempSvcNameCamelLower po.TempSvcNameCaseCamel, err error) {
	tx := repo.db.WithContext(ctx)

	err = tx.Where("id = ?", id).First(&TempSvcNameCamelLower).Error
	return
}

func (repo *TempSvcNameCamelLowerRepo) FindByQuery(ctx context.Context, filter map[string]interface{}, args ...interface{}) (res po.TempSvcNameCaseCamel, err error) {
	tx := repo.db.WithContext(ctx)

	dbQuery := tx.Model(&po.TempSvcNameCaseCamel{}).Where(filter)
	if len(args) >= 2 {
		dbQuery = dbQuery.Where(args[0], args[1:]...)
	} else if len(args) >= 1 {
		dbQuery = dbQuery.Where(args[0])
	}

	err = dbQuery.First(&res).Error
	return
}

func (repo *TempSvcNameCamelLowerRepo) FindList(ctx context.Context, filter map[string]interface{}, args ...interface{}) (total int64, res []po.TempSvcNameCaseCamel, err error) {
	tx := repo.db.WithContext(ctx)

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

	dbQuery := tx.Model(&po.TempSvcNameCaseCamel{}).Where(filter)

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
	tx := repo.db.WithContext(ctx)

	dbQuery := tx.Model(&po.TempSvcNameCaseCamel{}).Where(filter)
	if len(args) >= 2 {
		dbQuery = dbQuery.Where(args[0], args[1:]...)
	} else if len(args) >= 1 {
		dbQuery = dbQuery.Where(args[0])
	}

	err = dbQuery.Count(&total).Error
	return
}

func (repo *TempSvcNameCamelLowerRepo) Insert(ctx context.Context, TempSvcNameCamelLower po.TempSvcNameCaseCamel) (id int64, err error) {
	tx := repo.db.WithContext(ctx)

	err = tx.Create(&TempSvcNameCamelLower).Error
	if err != nil {
		return
	}

	id = TempSvcNameCamelLower.Id
	return
}

func (repo *TempSvcNameCamelLowerRepo) Update(ctx context.Context, id int64, TempSvcNameCamelLower po.TempSvcNameCaseCamel) (err error) {
	tx := repo.db.WithContext(ctx)

	err = tx.Model(&po.TempSvcNameCaseCamel{}).Where("id = ?", id).Updates(&TempSvcNameCamelLower).Error
	return
}

func (repo *TempSvcNameCamelLowerRepo) Delete(ctx context.Context, id int64) (err error) {
	tx := repo.db.WithContext(ctx)

	err = tx.Where("id = ?", id).Delete(&po.TempSvcNameCaseCamel{}).Error
	return
}


`
