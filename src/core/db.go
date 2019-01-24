package core

import (
	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Db struct {
	XormEngine *xorm.Engine
}

func CreateDbObject() (Db, error) {

	config := GetInstanceConfig()
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.User = config.Db.User
	mysqlConfig.DBName = config.Db.Name
	mysqlConfig.Passwd = config.Db.Password
	mysqlConfig.Params = make(map[string]string)
	mysqlConfig.Params["charset"] = config.Db.Charset
	mysqlConfig.Net = "tcp"
	mysqlConfig.Addr = config.Db.Host + ":" + config.Db.Port
	//helper.GetInstanceLog().Out(mysqlConfig.FormatDSN())
	engine, err := xorm.NewEngine("mysql", mysqlConfig.FormatDSN())
	ndb := Db{
		XormEngine: engine,
	}
	return ndb, err
}
