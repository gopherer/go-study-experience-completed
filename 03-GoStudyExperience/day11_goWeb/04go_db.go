package main

import (
	"database/sql"
	_ "github.com/mysql-master"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	if err != nil {
		panic(err.Error())
	}
}

//Sql未学待续。。。
