package main

import "fmt"

func main26() {
	m := map[int]string{1: "mike", 2: "go", 3: "hello"}
	//第一个返回值为key，第二个返回值为value，遍历结果是无序的
	for key, value := range m {
		fmt.Printf("%d =========== > %s\n", key, value)
	}
	//如何判断一个key值是否存在
	//第一个返回值为key所对应的value，第二个返回值为key是否存在的条件，存在为true，不存在为false
	value, ok := m[1]
	if ok == true {
		fmt.Println("m[1] = ", value)
	} else {
		fmt.Println("key不存在")
	}

	value, ok = m[0]
	if ok == true {
		fmt.Println("m[1] = ", value)
	} else {
		fmt.Println("key不存在")
	}
}
