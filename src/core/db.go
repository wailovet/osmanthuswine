package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-sql-driver/mysql"
	"github.com/wailovet/osmanthuswine/src/helper"
)

type Db struct {
	GormDB *gorm.DB
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
	helper.GetInstanceLog().Out(mysqlConfig.FormatDSN())
	db, err := gorm.Open("mysql", mysqlConfig.FormatDSN())
	ndb := Db{
		GormDB: db,
	}
	return ndb, err
}
