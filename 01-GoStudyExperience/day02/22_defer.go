package main

import "fmt"

func defer1() {
	defer fmt.Println("瞅你咋地")
	fmt.Println("你瞅啥")

}

func main22() {
	//defer延迟调用，main函数结束前调用 
	//defer fmt.Println("oooooooooooooooo")
	fmt.Println("sssssssssssssssss")
	defer fmt.Println("hhhhhhhhhhhhhhhhh")
	defer1()
}
