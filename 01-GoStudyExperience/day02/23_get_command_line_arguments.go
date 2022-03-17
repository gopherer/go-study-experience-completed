package main

import "fmt"
import "os"

func main23(){
	//接收用户传递的参数，都是以字符串方式传递
	list := os.Args
	n := len(list)
	fmt.Println("n = ",n)
	
	for i := 0; i < n; i++ {
		fmt.Printf("list[%d] = %s",i,list[i])
	}
	fmt.Println("")
	for i,data:= range list {
		fmt.Printf("list[%d] = %s",i,data)
	}
	
}
