package core

import (
	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
	"database/sql"
	"strconv"
)

type Db struct {
	XormEngine *xorm.Engine
	GormDB     *gorm.DB
}

var threadsConnectedNumConn *sql.DB

func GetThreadsConnectedNum() (int, error) {
	if threadsConnectedNumConn == nil {
		cc := GetInstanceConfig()
		mysqlConfig := mysql.NewConfig()
		mysqlConfig.User = cc.Db.User
		mysqlConfig.DBName = cc.Db.Name
		mysqlConfig.Passwd = cc.Db.Password
		mysqlConfig.Params = cc.Db.Params
		mysqlConfig.Net = "tcp"
		mysqlConfig.Addr = cc.Db.Host + ":" + cc.Db.Port
		threadsConnectedNumConn, _ = sql.Open("mysql", mysqlConfig.FormatDSN())
	}

	rows, err := threadsConnectedNumConn.Query("show status like 'Threads_connected';")
	if err != nil {
		return 0, err
	}
	cols, err := rows.Columns()
	if err != nil {
		return 0, err
	}
	buff := make([]interface{}, len(cols)) // 临时slice，用来通过类型检查
	data := make([]string, len(cols))      // 真正存放数据的slice
	for i, _ := range buff {
		buff[i] = &data[i] // 把两个slice关联起来
	}
	for rows.Next() {
		rows.Scan(buff...)
	}
	ii, _ := strconv.Atoi(data[1])
	return ii, nil
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
		db.DB().SetMaxOpenConns(config.Db.MaxOpenConn)
		db.SingularTable(true)
		instanceDb = &Db{
			GormDB: db,
		}
		return instanceDb, err
	}

	return instanceDb, nil
}
