package main

import (
	"encoding/json"
	"fmt"
)

//IT 成员变量首字母必须大写
type IT struct {
	Company  string   `json:"-"`        //此字段不会输出到屏幕
	Subjects []string `json:"subjects"` //二次编码 以subjects输出
	IsOk     bool     `json:",string"`  //以字符串格式输出
	Price    float64  `json:",string"`
}

func main11() {
	//定义一个结构体变量，同时初始化
	s := IT{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.888}

	//编码，根据内容生成josn文本
	//buf, err := json.Marshal(s)
	buf, err := json.MarshalIndent(s, "", " ") //格式化编码
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf))
}

// {
// 	"Company":"itcast",
// 	"Subjects":[
// 		"Go",
// 		"C++",
// 		"Python",
// 		"Test"
// 	],
// 	"IsOk":true,
// 	"Price":666.6666
// }
