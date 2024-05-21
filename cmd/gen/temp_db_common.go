package gen

const TempDBCommon = `
package driven

import (
	"context"
	"TempImportPkg/global"
	"TempImportPkg/infra/transaction"

	"gorm.io/gorm"
)

type DBOption func(*gorm.DB) *gorm.DB

func getDb(opts ...DBOption) *gorm.DB {
	db := global.DB
	for _, opt := range opts {
		db = opt(db)
	}
	return db
}

func getTx(ctx context.Context, opts ...DBOption) *gorm.DB {
	tx, ok := ctx.Value(transaction.DB).(*gorm.DB)
	if ok {
		for _, opt := range opts {
			tx = opt(tx)
		}
		return tx
	}
	return getDb(opts...)
}
`
