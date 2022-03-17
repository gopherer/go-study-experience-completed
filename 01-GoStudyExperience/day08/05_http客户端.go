package main

import (
	"fmt"
	"net/http"
)

func main05() {
	//resp, err := http.Get("http://www.baidu.com")
	resp, err := http.Get("http://127.0.0.1:8000/go")
	if err != nil {
		fmt.Println("http.Get err = ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status = ", resp.Status)
	fmt.Println("StatusCode = ", resp.StatusCode)
	fmt.Println("Header = ", resp.Header)
	fmt.Println("Body = ", resp.Body)

	buf := make([]byte, 4*1024)

	var tep string

	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.Read = ", err)
			break
		}
		tep += string(buf[:n])
	}
	fmt.Println("tep = ", tep)
}
