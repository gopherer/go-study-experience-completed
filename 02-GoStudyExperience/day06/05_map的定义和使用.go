package main

import "fmt"

func main05() {
	//var m map[int]string
	//fmt.Printf("%p\n", m) //地址为0 系统占用不允许读写

	//m[100] = "gogo"//err
	//fmt.Println(m)

	var m map[int]string = map[int]string{100: "hello", 120: "go"}
	fmt.Printf("%p\n", m)
	fmt.Println(m)

	m[200] = "hahaha"
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println(m[100])
}

func main0501() {
	m := make(map[string]int)
	fmt.Println(len(m))
	fmt.Printf("%p\n", m)
	//map 会自动扩容
	m["张山"] = 100
	m["李四"] = 200
	fmt.Println(len(m))
	fmt.Println(m)
	//重新赋值
	m["李四"] = 300
	fmt.Println(m)

}
