package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// GormMysql 初始化Mysql数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	//if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
	//	return nil
	//} else {
	//	sqlDB, _ := db.DB()
	//	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	//	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	//	return db
	//}
	fmt.Printf("%v\n", mysqlConfig)
	retryCnt := 1
	for {
		fmt.Printf("第%d次尝试连接\n", retryCnt)
		if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
			//global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
			//os.Exit(0)
			//return nil

		} else {
			sqlDB, _ := db.DB()
			sqlDB.SetMaxIdleConns(m.MaxIdleConns)
			sqlDB.SetMaxOpenConns(m.MaxOpenConns)
			return db
		}

		time.Sleep(5 * time.Second)
		retryCnt++
		if retryCnt >= 20 {
			return nil
		}
	}
}

// GormMysqlByConfig 初始化Mysql数据库用过传入配置
func GormMysqlByConfig(m config.DB) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
