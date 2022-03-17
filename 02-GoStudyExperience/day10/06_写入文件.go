package main

import (
	"fmt"
	"os"
)

func main06() {
	fp, err := os.Create("a.txt")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println(&err)
	defer func() {
		err := fp.Close()
		if err != nil {
			return
		}
	}()

	//将字符串写入文件中
	//\n不会换行 原因是在Windows文本文件中换行用\r\n来表示 在Linux中换行用\n表示
	n, err := fp.WriteString("hello go\r\n") //n 表示写入的字符的个数
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println(&err)
	fmt.Println(n)
	fp.WriteString("大家好")
}
func main0601() {
	fp, err := os.Create("a.txt")
	if err != nil {
		return
	}
	defer fp.Close()
	//slice := []byte{'h', 'e', 'l', 'l', 'o'}
	//fp.Write(slice)

	//slice := []byte("大家好")
	//fmt.Printf("%T\n", slice)
	//fp.Write(slice)

	count, err := fp.Write([]byte("下午好"))
	if err != nil {
		return
	} else {
		fmt.Println(count)
	}

}
func main0602() {
	//os.Open("文件路径") 只读方式打开文件
	//os.OpenFile("文件路径",打开方式,打开权限)
	//打开方式, os.O_RDONLY--只读  os.O_WRONLY--只写 os.O_RDWR 可读可写 os.O_APPEND --追加
	//打开权限，0-7 rwx 6(rw-) 7(rwx)
	fp, err := os.OpenFile("a.txt", os.O_WRONLY, 6)
	if err != nil {
		fmt.Println("未找到文件")
		return
	}
	defer fp.Close()
	//获取光标位置 获取从文件起始到结尾有多少个字符
	//count, err := fp.Seek(0, os.SEEK_END)
	//count, _ := fp.Seek(0, io.SeekEnd)
	//fp.WriteAt([]byte("hello world"), count)
	fp.WriteString("haha\r\n")
	fp.WriteAt([]byte("6666"), 6)

}
