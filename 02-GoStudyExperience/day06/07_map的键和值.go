package main

import "fmt"

func main07() {
	m := make(map[string][3]int)

	m["张三"] = [3]int{100, 200, 400}
	m["李四"] = [3]int{10, 20, 30}
	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("姓名：%s 语文：%d 数学：%d 英语：%d \n", k, v[0], v[1], v[2])

		for i := 0; i < len(v); i++ {
			fmt.Println(v[i])
		}
	}
}

func main0701() {
	m := make(map[int]string)

	m[100] = "白起"
	m[200] = "嬴政"

	value, ok := m[100]

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("未找到")
	}

	//未找到key 打印value类型默认的值
	fmt.Println(m[300])
}

func main0702() {
	m := make(map[int]string)

	m[100] = "白起"
	m[200] = "嬴政"

	fmt.Println(m)
	delete(m, 100)
	fmt.Println(m)
	delete(m, 200)
	fmt.Println(m)
	//删除key值 如果key不存在 不会报错
	delete(m, 400)
	fmt.Println(m)
}
