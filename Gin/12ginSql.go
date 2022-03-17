package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	//链接数据库
	conStr := "root:root@tcp(127.0.0.1:3306)/ginsql"
	db, _ := sql.Open("mysql", conStr)
	//创建数据库表
	//person:id, name, age
	_, err := db.Exec("create table if not exists person(" +
		"id int auto_increment primary key," +
		"name varchar(12) not null," +
		"age int default 1" +
		");")
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("数据库创建成功")
	}
	//数据插入到数据表
	if _, err := db.Exec("insert into person(name,age)"+
		"values(?,?),(?,?)", "Davie", 18, "张三", 20); err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("数据插入成功")
	}
	//查询数据库
	rows, err := db.Query("select * from person")
	if err != nil {
		log.Fatal(err.Error())
	}
	for rows.Next() {
		p := new(person)
		err := rows.Scan(&p.Id, &p.Name, &p.Age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(p)
	}
	//flag:
	//	if rows.Next() {
	//		p := new(person)
	//		err := rows.Scan(&p.Id, &p.Name, &p.Age)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		fmt.Println(p)
	//		goto flag
	//	}
}
