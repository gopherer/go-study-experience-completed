package main

import "fmt"

func main08() {
	var ch byte
	ch = 97
	fmt.Println("ch = ", ch)
	//格式化输出，%c以字符方式打印，%d以整型方式打印
	fmt.Printf("%c %d\n", ch, ch)

	ch = 'a' //字符，单引号
	fmt.Printf("%c %d\n", ch, ch)

	//大写转小写，小写转大写，大小写相差32，小写大
	fmt.Printf("大写：%d, 小写: %d\n", 'A', 'a')
	fmt.Printf("大写转小写:%c\n", 'A'+32)
	fmt.Printf("小写转大写:%c\n", 'a'-32)

	//'\'以反斜杠开头的字符是转义字符，'\n'代表换行
	fmt.Printf("hello go%c", '\n') //10 = '\n'
	fmt.Printf("hello itcast")
}
