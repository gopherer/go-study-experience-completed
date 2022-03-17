package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main06() {
	var start, end int
	fmt.Println("起始页(>=1): ")
	fmt.Scan(&start)
	fmt.Println("终止页(>=起始页): ")
	fmt.Scan(&end)

	DoWork(start, end)

}

//DoWork 开始工作
func DoWork(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d 的页面\n", start, end)

	//明确目标(要知道你准备在哪些范围或网站去搜索)
	//https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		fmt.Println("url = ", url)

		//将网页所有的内容爬取下来
		result, err := HTTPGet(url)
		if err != nil {
			fmt.Println("HTTPGet err = ", err)
			continue
		}

		//把内容写入文件
		fileName := "第" + strconv.Itoa(i) + "页.html"
		f, err := os.Create(fileName)
		if err != nil {
			fmt.Println("os.Creat err = ", err)
			continue
		}

		f.WriteString(result)

		f.Close()
	}
}

//HTTPGet 爬取网页信息
func HTTPGet(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}

	defer resp.Body.Close()

	//读取网页信息Body的内容
	buf := make([]byte, 4*1024)

	for {
		n, err := resp.Body.Read(buf)
		if n == 0 { //读取结束，或出问题
			fmt.Println("resp.Body.Read err = ", err)
			break
		}
		result += string(buf[:n])
	}
	return
}
