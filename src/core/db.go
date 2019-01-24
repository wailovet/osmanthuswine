package core

import (
	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
	"github.com/go-errors/errors"
)

type Db struct {
	XormEngine *xorm.Engine
	GormDB     *gorm.DB
}

const DbTypeXorm = 0
const DbTypeGorm = 1

func CreateDbObject(dbtype int) (*Db, error) {
	config := GetInstanceConfig()
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.User = config.Db.User
	mysqlConfig.DBName = config.Db.Name
	mysqlConfig.Passwd = config.Db.Password
	mysqlConfig.Params = config.Db.Params
	mysqlConfig.Net = "tcp"
	mysqlConfig.Addr = config.Db.Host + ":" + config.Db.Port

	//helper.GetInstanceLog().Out(mysqlConfig.FormatDSN())

	if dbtype == DbTypeXorm {
		engine, err := xorm.NewEngine("mysql", mysqlConfig.FormatDSN())
		ndb := Db{
			XormEngine: engine,
		}
		return &ndb, err
	}
	if dbtype == DbTypeGorm {
		db, err := gorm.Open("mysql", mysqlConfig.FormatDSN())
		ndb := Db{
			GormDB: db,
		}
		return &ndb, err
	}
	return nil, errors.New("error orm type")
}
