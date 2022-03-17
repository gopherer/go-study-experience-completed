package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	H = 3
	L = 3
)

type sparse struct {
	row   int
	col   int
	value int
}

func main() {
	var sparseArr [H][L]int
	sparseArr[1][1] = 1 //白子
	sparseArr[2][2] = 2 //黑子

	//遍历sparseArr
	fmt.Println("原始数组")
	for _, v1 := range sparseArr {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	//记录稀疏数组第一行
	var sparseSlice []sparse
	sparseSlice = append(sparseSlice, sparse{
		row:   len(sparseArr),
		col:   len(sparseArr[0]),
		value: 0,
	})
	//给稀疏数组赋值
	for i, v1 := range sparseArr {
		for j, v2 := range v1 {
			if v2 != 0 {
				sparseSlice = append(sparseSlice, sparse{
					row:   i,
					col:   j,
					value: v2,
				})
			}
		}
	}
	fmt.Println("稀疏数组")
	for _, v1 := range sparseSlice {
		fmt.Println(v1.row, v1.col, v1.value)
	}
	//将稀疏数组写入文件
	writeFile(sparseSlice)
	//从文件读取稀疏数组
	test, err := readFile()
	if err != nil {
		fmt.Println("read file err", err)
	}
	fmt.Println(test)
	var newArr [H][L]int
	for i, v := range test {
		if i != 0 {
			newArr[v.row][v.col] = v.value
		}
	}
	for _, v1 := range newArr {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

}
func writeFile(slice []sparse) {
	file, err := os.Create("C:\\GoCode\\src\\03_gocode\\day09\\sparseData.txt")
	if err != nil {
		fmt.Println("file creat err", err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("file close err", err)
		}
	}()
	for _, value := range slice {
		temp := strconv.Itoa(value.row) + " " + strconv.Itoa(value.col) + " " + strconv.Itoa(value.value) + string('\n')
		_, err = file.WriteString(temp)
		if err != nil {
			fmt.Println("file writeString err", err)
		}
	}
}
func readFile() (test []sparse, err error) {
	file, err := os.Open("C:\\GoCode\\src\\03_gocode\\day09\\sparseData.txt")
	if err != nil {
		fmt.Println("file open err", err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("file close err", err)
		}
	}()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			//fmt.Println("readSlice err", err)
			break
		}
		//fmt.Println(str)
		newStr := strings.Trim(str, "\n")
		strSlice := strings.Split(newStr, " ")

		//fmt.Println(strSlice)
		i, _ := strconv.Atoi(strSlice[0])
		j, _ := strconv.Atoi(strSlice[1])
		k, err := strconv.Atoi(strSlice[2])
		//fmt.Println(err)
		test = append(test, sparse{
			row:   i,
			col:   j,
			value: k,
		})

	}
	//fmt.Println(test)
	return
}
