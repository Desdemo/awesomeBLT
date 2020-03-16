package model

import (
	_ "github.com/go-sql-driver/mysql"
	"log"

	"xorm.io/xorm"
)

var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("mysql", "saler:saler@(192.168.0.36)/blt113?charset=utf8mb4")
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	Engine.ShowSQL(true)

}
