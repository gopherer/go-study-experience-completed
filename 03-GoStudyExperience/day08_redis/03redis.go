package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main03() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis dial err = ", err)
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("conn.Close err", err)
		}
	}()
	_, err = conn.Do("HMSet", "user02", "name", "mary", "addr", "beijing")
	if err != nil {
		fmt.Println("conn err = ", err)
	}
	r, err := redis.Strings(conn.Do("HmGet", "user02", "name", "addr"))
	if err != nil {
		fmt.Println("redis.Strings conn do err = ", err)
		return
	}
	for i, v := range r {
		fmt.Printf("r[%d] = %s\n", i, v)
	}
}
