package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CountCharacter struct {
	charCount    int
	spaceCount   int
	numCount     int
	chineseCount int
	otherCount   int
}

func main06() {
	var count CountCharacter
	file, err := os.Open("C:/GoCode/src/03_gocode/day05_file/test.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Print("统计结束")
				break
			}
			panic(err)
		}
		newStr := []rune(str)
		for _, v := range newStr {
			switch {
			case v > 'a' && v < 'z':
				fallthrough
			case v > 'A' && v < 'Z':
				count.charCount++
			case string(v) == " " || string(v) == "\t":
				count.spaceCount++
			case v > '0' && v < '9':
				count.numCount++
			case v > 0x4E00 && v < 0x9FA5:
				count.chineseCount++
			default:
				count.otherCount++
			}
		}
	}
	fmt.Printf("char = %v space = %v num = %v chinese = %v other = %v",
		count.charCount, count.spaceCount, count.numCount, count.chineseCount, count.otherCount)

}
