package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//WriteFile 打开文件
func WriteFile(path string) {
	//打开文件，新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close() //延迟调用 在函数结束前关闭文件

	var buf string

	for i := 0; i < 10; i++ {
		//format string ====== "i = %d\n"
		//buf = fmt.Sprintf("i = %d\n",i)
		buf = fmt.Sprintln("i = ", i)
		_, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
	}

}

//ReadFile 读取文件
func ReadFile(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//关闭文件
	defer f.Close()

	buf := make([]byte, 1024*2) //2k大小

	//n表示从文件读取内容的长度
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF { //文件出错，同时没有到文件结尾
		fmt.Println("err = ", err)
		return
	}
	fmt.Println(string(buf[:n])) //切片从下标0取n个元素
}

//ReadFileline 读取文件
func ReadFileline(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//关闭文件
	defer f.Close()

	//新建一个缓冲区,先把内容先放入缓冲区
	r := bufio.NewReader(f)

	for {
		//遇到'\n'结束读取，但是'\n'也会读取进去
		buf ,err := r.ReadBytes('\n')

		if err != nil{
			if err == io.EOF{//文件已到结尾
				break
			}
			fmt.Println("err = ",err)
			return 
		}
		fmt.Printf("buf = !%s!\n",string(buf))

	}
}
func main16() {
	path := "./16_demo.txt"
	WriteFile(path)
	//ReadFile(path)
	ReadFileline(path)
}
