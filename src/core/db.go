package core

import (
	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
	"github.com/go-errors/errors"
	"database/sql"
	"strconv"
	"time"
)

type Db struct {
	XormEngine *xorm.Engine
	GormDB     *gorm.DB
}

var ThreadsConnectedNum  = 0
func GetThreadsConnectedNum() int {

	cc := GetInstanceConfig()
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.User = cc.Db.User
	mysqlConfig.DBName = cc.Db.Name
	mysqlConfig.Passwd = cc.Db.Password
	mysqlConfig.Params = cc.Db.Params
	mysqlConfig.Net = "tcp"
	mysqlConfig.Addr = cc.Db.Host + ":" + cc.Db.Port
	threadsConnectedNumConn, _ := sql.Open("mysql", mysqlConfig.FormatDSN())

	rows, _ := threadsConnectedNumConn.Query("show status like 'Threads_connected';")

	defer  rows.Close()
	defer  threadsConnectedNumConn.Close()

	cols, _ := rows.Columns()
	buff := make([]interface{}, len(cols)) // 临时slice，用来通过类型检查
	data := make([]string, len(cols))      // 真正存放数据的slice
	for i, _ := range buff {
		buff[i] = &data[i] // 把两个slice关联起来
	}
	for rows.Next() {
		rows.Scan(buff...)
	}
	ii, _ := strconv.Atoi(data[1])
	return ii
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

	for GetThreadsConnectedNum() > config.Db.MaxOpenConn {
		time.Sleep(time.Second)
	}

	if dbtype == DbTypeXorm {
		engine, err := xorm.NewEngine("mysql", mysqlConfig.FormatDSN())
		engine.SetMaxOpenConns(config.Db.MaxOpenConn)
		ndb := Db{
			XormEngine: engine,
		}
		return &ndb, err
	}
	if dbtype == DbTypeGorm {
		db, err := gorm.Open("mysql", mysqlConfig.FormatDSN())
		db.DB().SetMaxOpenConns(config.Db.MaxOpenConn)
		ndb := Db{
			GormDB: db,
		}
		return &ndb, err
	}
	return nil, errors.New("error orm type")
}
