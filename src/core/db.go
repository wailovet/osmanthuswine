package core

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var instanceDb *gorm.DB

func GetDb() (*gorm.DB, error) {
	if instanceDb == nil {
		config := GetInstanceConfig()
		mysqlConfig := mysql.NewConfig()
		mysqlConfig.User = config.Db.User
		mysqlConfig.DBName = config.Db.Name
		mysqlConfig.Passwd = config.Db.Password
		mysqlConfig.Params = config.Db.Params
		mysqlConfig.Net = "tcp"
		mysqlConfig.Addr = config.Db.Host + ":" + config.Db.Port

		db, err := gorm.Open("mysql", mysqlConfig.FormatDSN())
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return config.Db.Prefix + defaultTableName
		}
		db.DB().SetMaxOpenConns(config.Db.MaxOpenConn)
		db.SingularTable(true)
		return db, err
	}
	return instanceDb, nil
}

func init() {
	go func() {
		for ; ; {
			if instanceDb != nil {
				err := instanceDb.DB().Ping()
				if err != nil {
					println(err.Error())
					instanceDb.Close()
					instanceDb = nil
				}
			}
			time.Sleep(time.Second)
		}
	}()
}
