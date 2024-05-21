package gen

const TempTransaction = `
package transaction

import (
	"context"
	"TempImportPkg/global"

	"gorm.io/gorm"
)

type DBContext string

const (
	DB DBContext = "db"
)

func GetTxAndContext() (tx *gorm.DB, ctx context.Context) {
	tx = global.DB.Begin()
	ctx = context.WithValue(context.Background(), DB, tx)
	return
}
`
