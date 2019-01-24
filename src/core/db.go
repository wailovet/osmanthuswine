package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-sql-driver/mysql"
)

type Db struct {
	gorm.DB
}

func CreateDbObject() (Db, error) {

	config := GetInstanceConfig()
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.User = config.Db.User
	mysqlConfig.DBName = config.Db.Name
	mysqlConfig.Passwd = config.Db.Password
	mysqlConfig.Params = make(map[string]string)
	mysqlConfig.Params["charset"] = config.Db.Charset
	mysqlConfig.Addr = config.Db.Host + ":" + config.Db.Port

	db, err := gorm.Open("mysql", mysqlConfig.FormatDSN())
	var sdb interface{}
	sdb = *db
	ndb := sdb.(Db)
	return ndb, err
}
