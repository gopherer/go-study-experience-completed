package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string `json:"monster_name"` //反射机制
	Age  int    `json:"age"`
	Addr string
}

func marshalStruct() []byte {
	var monster = Monster{
		Name: "白骨精",
		Age:  400,
		Addr: "未知",
	}
	b, err := json.Marshal(&monster) //json调用monster结构体，若monster内字段是小写则无法调用
	if err != nil {
		fmt.Print("monster序列化失败err = ", err)
	}
	fmt.Printf("monster序列化后 = %v\n", string(b))
	return b
}
func marshalMap() {
	var m = make(map[string]interface{})
	m["name"] = "孙悟空"
	m["age"] = 500
	m["addr"] = "花果山"
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Print("map序列化失败err = ", err)
	}
	fmt.Printf("map序列化后 = %v\n", string(b))
}
func marshalSlice() {
	var slice []map[string]interface{}
	m := make(map[string]interface{})
	m["name"] = "孙悟空"
	m["age"] = 500
	m["addr"] = "花果山"
	slice = append(slice, m)
	b, err := json.Marshal(slice)
	if err != nil {
		fmt.Print("slice序列化失败err = ", err)
	}
	fmt.Printf("slice序列化后 = %v\n", string(b))
}
func marshalFloat() {
	var num = 9.99
	b, err := json.Marshal(num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("slice序列化后 = %v\n", string(b))
}
func main09() {
	marshalStruct()
	marshalMap()
	marshalSlice()
	marshalFloat()
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
func unmarshalStruct() {
	var m Monster
	b := marshalStruct()
	err := json.Unmarshal(b, &m) //注意要传入地址
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m, m.Name)
}
func unmarshalMap() {
	var m map[string]interface{}
	var str = "{\"addr\":\"花果山\",\"age\":500,\"name\":\"孙悟空\"}" //注意{}
	err := json.Unmarshal([]byte(str), &m)                      //注意&  unmarshal会在底层make
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}
func unmarshalSlice() {
	var s []map[string]interface{}
	var str = `[{"addr":"花果山","age":500,"name":"孙悟空"}]`
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("s = ", s)
}
