package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main02() {
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
	_, err = conn.Do("HSet", "user01", "name", "john")
	if err != nil {
		fmt.Println("conn do err", err)
		return
	}
	r, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("redis.String conn do err ", err)
		return
	}
	fmt.Printf("r = %v\n", r)

}
