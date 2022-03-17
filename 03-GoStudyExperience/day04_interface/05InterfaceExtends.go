package main

import "fmt"

type Bird interface {
	Flying()
}
type Fish interface {
	Swimming()
}
type Monkey struct {
	Name string
}
type LittleMonkey struct {
	Monkey
}

func (monkey Monkey) climbing() {
	fmt.Println(monkey.Name, "会爬树....")
}
func (monkey LittleMonkey) Flying() {
	fmt.Println(monkey.Name, "会飞...")
}
func (monkey LittleMonkey) Swimming() {
	fmt.Println(monkey.Name, "会游泳...")
}

func main05() {
	var monkey = LittleMonkey{Monkey{Name: "悟空"}}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
