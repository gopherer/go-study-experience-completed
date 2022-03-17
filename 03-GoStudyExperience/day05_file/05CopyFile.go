package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(desFilePath string, srcFilePath string) (int64, error) {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Print("file open err")
		panic(err)
	}
	defer func() {
		err = srcFile.Close()
		if err != nil {
			fmt.Print("file close err")
		}
	}()
	reader := bufio.NewReader(srcFile)

	desFile, err := os.OpenFile(desFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Print("file open err")
		panic(err)
	}
	defer func() {
		err = desFile.Close()
		if err != nil {
			fmt.Print("file close err")
		}
	}()
	writer := bufio.NewWriter(desFile)
	return io.Copy(writer, reader)

}
func main05() {
	var srcFilePath = "C:/GoCode/src/03_gocode/day05_file/photo.png"
	var desFilePath = "C:/GoCode/src/03_gocode/day05_file/test.png"
	_, err := CopyFile(desFilePath, srcFilePath)
	if err == nil {
		fmt.Print("file copy complete")
	}
}
