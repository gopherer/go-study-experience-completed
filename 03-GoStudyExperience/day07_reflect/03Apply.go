package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func (m Monster) Get() {
	fmt.Println(m.Name, m.Age, m.Gender)
}
func (m Monster) Set(name string, age int, gender string) {
	m.Name = name
	m.Age = age
	m.Gender = gender
}
func Reflect(b interface{}) {
	rType := reflect.TypeOf(b)
	rValue := reflect.ValueOf(b)
	kd := rValue.Kind()
	if kd != reflect.Struct {
		fmt.Println("unexpected type")
		return
	}
	numType := rValue.NumField()
	fmt.Printf("%v共有%v个字段\n", kd, numType)
	for i := 0; i < numType; i++ {
		fmt.Printf("fidld[%v] = %v\n", i, rValue.Field(i))
		//获取json值
		tag := rType.Field(i).Tag.Get("json")
		if tag != "" {
			fmt.Printf("%v 第%v 的 tag = %v\n", kd, i, tag)
		}
	}
	numOfMethod := rValue.NumMethod()
	fmt.Printf("struct 共有%v方法\n", numOfMethod)
	rValue.Method(0).Call(nil)

}

func main03() {
	monster := Monster{
		Name:   "jack",
		Age:    18,
		Gender: "man",
	}
	Reflect(monster)
}
