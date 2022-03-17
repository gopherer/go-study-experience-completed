package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
		TestOnBorrow:    nil,
		MaxIdle:         8,   //最大空闲链接
		MaxActive:       0,   //数据库最大连接数，0表示没有限制
		IdleTimeout:     100, //最大空闲时间
		Wait:            false,
		MaxConnLifetime: 0,
	}
}
func main04() {
	conn := pool.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("conn Close err=", err)
		}
	}()
	_, err := conn.Do("Set", "name", "john")
	if err != nil {
		fmt.Println("conn do err ", err)
		return
	}
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("redis.String conn do err =", err)
		return
	}
	fmt.Println("r = ", r)
}
