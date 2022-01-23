package database

import (
	"fmt"
	"time"

	"github.com/Yan-Bin-Lin/DCreater/app/setting"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	db *xorm.Engine
)

func init() {
	var (
		err        error
		connectStr string
		dbName     string
	)

	if setting.Servers["main"].RunMode == "debug" {
		dbName = "test"
	} else {
		dbName = "main"
	}
	//connectStr = fmt.Sprintf("%s:%s@/%s?%s", setting.DBs[dbName].User, setting.DBs[dbName].Password, setting.DBs[dbName].Name, setting.DBs[dbName].Param)
	connectStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", setting.DBs[dbName].User, setting.DBs[dbName].Password,
		setting.DBs[dbName].Host, setting.DBs[dbName].Port, setting.DBs[dbName].Name, setting.DBs[dbName].Param)

	db, err = xorm.NewEngine(setting.DBs["main"].Driver, connectStr)
	if err != nil {
		panic(err)
	}

	// optimize option
	db.SetMaxOpenConns(setting.DBs["main"].Option["SetMaxOpenConnects"])
	db.SetMaxIdleConns(setting.DBs["main"].Option["SetMaxIdleConnects"])
	db.SetConnMaxLifetime(time.Duration(setting.DBs["main"].Option["SetConnMaxLifetime"]) * time.Second)

	if setting.Servers["main"].RunMode == "debug" {
		db.ShowSQL(true)
		db.Logger().SetLevel(0) //core.LOG_DEBUG)
	}
}
