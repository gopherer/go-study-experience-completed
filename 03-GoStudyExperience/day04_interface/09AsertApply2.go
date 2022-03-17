package main

import "fmt"

type Giant struct {
}

func TypeJudge(item ...interface{}) {
	for index, value := range item {
		switch value.(type) {
		case int, int8, int16, int32, int64:
			fmt.Printf("第%v的元素数据类型为%T\n", index+1, value)
		case float64:
			fmt.Printf("第%v的元素数据类型为%T\n", index+1, value)
		case float32:
			fmt.Printf("第%v的元素数据类型为%T\n", index+1, value)
		case string:
			fmt.Printf("第%v的元素数据类型为%T\n", index+1, value)
		case Giant:
			fmt.Printf("第%v的元素数据类型为%T\n", index+1, value)
		case *Giant:
			fmt.Printf("第%v的元素数据类型为%T\n", index+1, value)
		}
	}
}
func main09() {
	var n = 10
	var n1 = 2.33
	var n2 float32 = 3.99
	var name = "tom"
	giant := Giant{}
	giant1 := &Giant{}
	TypeJudge(n, n1, n2, name, giant, giant1)
}
