package core

import (
	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
)

type Db struct {
	XormEngine *xorm.Engine
	GormDB     *gorm.DB
}

var instanceDb *Db

func CreateDbObject() (*Db, error) {
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
		instanceDb = &Db{
			GormDB: db,
		}
		return instanceDb, err
	}
	return instanceDb, nil
}
