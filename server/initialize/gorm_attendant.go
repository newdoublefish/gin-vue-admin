package initialize

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func AttendantDb() *gorm.DB {
	dsn := os.Getenv("ATTENDANT_DB")
	if dsn == ""{
		dsn = "sqlserver://sa:dns.com.hk.2011@10.0.0.9?database=powinkq&encrypt=disable"
	}

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err == nil {
		return db
	}
	fmt.Printf(err.Error())
	return nil
}