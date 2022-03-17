package main

import (
	"os"
	"os/exec"
)

//清屏
func Clear() error {
	cmd := exec.Command("cmd.exe", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	return err
}

//切片倒置
func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
