package main

import "fmt"

type student struct {
	name  string
	age   int
	score int
}

func main01() {
	m := make(map[int]student)
	m[100] = student{
		name:  "张三",
		age:   18,
		score: 59,
	}
	m[200] = student{
		name:  "李四",
		age:   20,
		score: 100,
	}
	//m[100].age = 100 //err 可以使用指针解决 value的地址不可寻 或者 更新结构体

	delete(m, 200)
	fmt.Println(m)
}

func main0101() {
	m := make(map[int][]student)

	m[100] = []student{{
		name:  "李四",
		age:   20,
		score: 100,
	}, {
		name:  "王五",
		age:   30,
		score: 80,
	},
	}

	m[100] = append(m[100], student{
		name:  "张三",
		age:   30,
		score: 70,
	})
	//m[100][1].age = 300 ok

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(m)
}
