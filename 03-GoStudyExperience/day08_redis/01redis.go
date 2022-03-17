package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main01() {
	//通过go 向redis 写入读取数据
	//1、链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis dial err ", err)
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("conn.Close err", err)
		}
	}()
	//2、通过go向redis写入数据
	_, err = conn.Do("Set", "name", "mary")
	if err != nil {
		fmt.Println("conn do err ", err)
		return
	}
	//3、向redis读取数据
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("redis read err", err)
		return
	}
	fmt.Println("r = ", r)
}
