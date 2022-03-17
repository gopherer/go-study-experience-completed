package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main07() {
	//以只读方式打开
	fp, err := os.Open("a.txt")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer fp.Close()
	buf := make([]byte, 3)
	//	n, _ := fp.Read(buf)
	//	fmt.Print(string(buf[:n]))
	//	n, _ = fp.Read(buf)
	//	fmt.Print(string(buf[:n]))
	//	n, _ = fp.Read(buf)
	//	fmt.Print(string(buf[:n]))
	for {
		n, err := fp.Read(buf)
		if err == io.EOF {
			//io.EOF 表示文件的结尾
			//当读到文件结尾时 返回值errors.New("EOF")
			fmt.Println("err = ", err)
			break
		}
		fmt.Print(string(buf[:n]))
	}
}
func main0701() {
	fp, err := os.OpenFile("a.txt", os.O_RDONLY, 6)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer fp.Close()
	//创建文件缓冲区
	r := bufio.NewReader(fp)
	slice, _ := r.ReadBytes('\n')
	fmt.Print(string(slice))
	slice, _ = r.ReadBytes('\n')
	fmt.Print(string(slice))

	for {
		slice, err := r.ReadBytes('\n')
		//先打印后判断
		fmt.Println(string(slice))
		if err == io.EOF {
			break
		}
	}
}
