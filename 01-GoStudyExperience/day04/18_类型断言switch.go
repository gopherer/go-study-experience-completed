package main

import "fmt"

//Girl 内容
type Girl struct {
	name string
	id   int
}

func main18() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello go"
	i[2] = Girl{"yoyo", 666}

	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d] 类型int 内容%d\n", index, value)
		case string:
			fmt.Printf("x[%d] 类型string 内容%s\n", index, value)
		case Girl:
			fmt.Printf("x[%d] 类型Giel 内容name = %s id = %d\n", index, value.name, value.id)
		}
	}
}
