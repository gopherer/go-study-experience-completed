package main

import (
	"encoding/json"
	"fmt"
)

func main12() {
	//创建一个map
	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"Go", "C++", "Python", "Test"}
	m["isok"] = true
	m["price"] = 6666.666

	//编码成json
	// result, err := json.Marshal(m)
	result, err := json.MarshalIndent(m,""," ")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//fmt.Println("err = ", err)
	fmt.Println("result = ", string(result))
}
