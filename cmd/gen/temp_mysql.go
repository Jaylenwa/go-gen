package gen

const TempMysql = `
package mysql

import (
	"fmt"
	"TempImportPkg/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewDB() *gorm.DB {
	return initDB()
}

func initDB() *gorm.DB {
	c := global.GConfig.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", c.Username, c.Password, c.DbHost, c.DbPort, c.DbName, c.Charset)

	cfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名
			SingularTable: true,
		},
		// sql打印
		Logger: logger.Default.LogMode(logger.Info),
	}
	db, err := gorm.Open(mysql.Open(dsn), cfg)
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	//连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	sqlDB.SetMaxIdleConns(c.MaxOpenConns)
	//设置数据库连接池最大连接数
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime) * time.Second)

	// 自动创建表
	err = db.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
`
