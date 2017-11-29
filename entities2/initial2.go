package entities2

import (
	_ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3308)/test?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
}
