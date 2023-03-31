package service

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	engine *xorm.Engine
}

var DDB *DB

func init() {
	engine, err := xorm.NewEngine("sqlite3", "/Users/wangrui/go/src/github.com/wrpromail/annotate-helper/pkg/dao/object_storage.db")
	if err != nil {
		panic(err)
	}
	DDB = &DB{engine: engine}
}
