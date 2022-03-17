package main

import (
	"encoding/json"
	"fmt"
)

//如果struct结构体中的字段 首字母没有大写 则自能在本包使用 无法被外界包所调用 例如encoding/json
type Monster struct { //struct ---tag
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr"`
}

func main06() {
	var monster Monster = Monster{"牛魔王", 500, "山洞"}
	fmt.Println("monster = ", monster)

	JsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("JsonStr = %s\n", string(JsonStr))
}
