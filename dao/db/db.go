package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

// 数据库初始化
func Init() (err error) {
	dns := "go_user:Cernet@goadmin@tcp(10.227.10.78:3399)/blog?charset=utf8&parseTime=True&loc=Local&timeout=1000ms"
	DB, err = sqlx.Connect("mysql", dns)
	if err != nil {
		log.Fatal("Connect database failed.", err)
		return
	}
	//数据库参数设置
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(10)
	return
}
