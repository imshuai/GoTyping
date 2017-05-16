package main

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

func DatabaseInit() {
	db := xorm.NewEngine("mysql", "ttt")
	db.SetMapper(core.GonicMapper{})
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(100)
	db.Sync2(&Artical{}, &Comment{})
}
