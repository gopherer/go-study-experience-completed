package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	interfaceChan := make(chan interface{}, 3)

	interfaceChan <- 10
	interfaceChan <- "hello"
	interfaceChan <- Cat{
		Name: "小花",
		Age:  2,
	}
	<-interfaceChan
	<-interfaceChan
	NewCat := <-interfaceChan
	fmt.Printf("NewCat type %T content %v\n", NewCat, NewCat)
	//如何直接调用NewCat.Name 会报错
	//fmt.Println(NewCat.Name)  NewCat.Name undefined (type interface {} is interface with no methods)
	//需要先类型断言
	cat := NewCat.(Cat)
	fmt.Printf("cat type %T cat.Name = %v", cat, cat.Name)
}
