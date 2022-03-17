package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main09() {
	fp, err := os.Open("a.txt")
	if err != nil {
		return
	}
	defer fp.Close()

	//创建缓冲区
	r := bufio.NewReader(fp)
	dict := make(map[string]string)
	for {
		word, _ := r.ReadBytes('\n')
		trans, err := r.ReadBytes('\n')
		dict[string(word[1:len(word)-2])] = string(trans[1 : len(trans)-2])
		if err == io.EOF {
			break
		}
	}
	fmt.Println("请输入要查找的单词：")
	var word string
	fmt.Scan(&word)
	value, ok := dict[word]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("未找到该单词")
	}
}
