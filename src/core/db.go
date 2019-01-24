package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Db struct {
	gorm.DB
}

var dbName = "test"
var dbPassword = "root"
var dbUser = "root"
var dbCharset = "utf8"

func CreateDbObject() (Db, error) {
	db, err := gorm.Open("mysql", dbUser+":"+dbPassword+"@/"+dbName+"?charset="+dbCharset+"&parseTime=True&loc=Local")
	var sdb interface{}
	sdb = *db
	ndb := sdb.(Db)
	return ndb, err
}
