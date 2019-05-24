package helper

import (
	"errors"
	"github.com/sonyarouje/simdb/db"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var dbName string

func getDriver() (*db.Driver, error) {
	var err error
	if dbName == "" {
		err = initDb()
		if err != nil {
			return nil, err
		}
	}
	driver, err := db.New(dbName)
	return driver, err
}

func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

func initDb() error {
	dir, _ := getCurrentPath()
	dbName = dir + "db"
	err := os.MkdirAll(dbName, os.ModeDir)
	if err != nil {
		log.Println(dbName, "没有权限")
	}
	return err
}

func JsonDbDriver() *db.Driver {
	mdb, err := getDriver()
	if err != nil {
		panic(err.Error())
	}
	return mdb
}

func JsonDbDriverOpen(parent db.Entity) *db.Driver {
	mdb, err := getDriver()
	if err != nil {
		panic(err.Error())
	}
	return mdb.Open(parent)
}

type JsonDb struct {
	UUID string `json:"uuid"`
}

func (that *JsonDb) ID() (jsonField string, value interface{}) {
	return "uuid", that.UUID
}
