package main

import (
	"encoding/json"
	"fmt"
)

//IT1 内容
type IT1 struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"` //二次编码 以subjects输出
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main13() {

	jsonbuf := `
	{
		"company":"itcast",
		"subjects":[
			"Go",
			"C++",
			"Python",
			"Test"
		],
		"isok":true,
		"price":666.6666
	}`
	//定义一个结构体变量
	var tmp IT1
	err := json.Unmarshal([]byte(jsonbuf), &tmp)//第二个参数为地址
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("tmp = ", tmp)
	fmt.Printf("tmp = %+v\n", tmp)

}
