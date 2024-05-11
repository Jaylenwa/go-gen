package gen

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"

	"github.com/gogf/gf-cli/v2/library/mlog"
)

func GetTableCol(table string) []TableColumn {
	var tc []TableColumn
	err := DB.Raw("SHOW FULL COLUMNS FROM `" + table + "`").Scan(&tc).Error
	if err != nil {
		panic(err)
	}

	return tc
}

func InitDB(host, port, user, password, dbName string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, port, dbName, "utf8mb4")
	cfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	}
	db, err := gorm.Open(mysql.Open(dsn), cfg)
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(50) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	sqlDB.SetMaxOpenConns(10) //设置数据库连接池最大连接数
	sqlDB.SetConnMaxLifetime(time.Duration(500) * time.Second)

	mlog.Print("db conn: ", dsn)

	DB = db
}
