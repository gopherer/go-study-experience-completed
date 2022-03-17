package main

import (
	"fmt"
	"strconv"
	"time"
)

func main06() {
	//获取当前时间
	t := time.Now()
	fmt.Printf("t = %v t type is %T\n", t, t)

	fmt.Printf("year = %v\n", t.Year())
	fmt.Printf("year type is %T\n", t.Year())
	fmt.Println("month = ", t.Month())
	fmt.Println("day = ", t.Day())
	fmt.Println("hour = ", t.Hour())
	fmt.Println("minute = ", t.Minute())
	fmt.Println("second = ", t.Second())

	timestr := fmt.Sprintf("current time %v-%d-%v %v:%v:%v ", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println(timestr)

	//格式化日期
	//fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Printf(t.Format("2006-01-02 15:04:05"))
	fmt.Println(t.Format("year = 2006"))
	fmt.Println(t.Format("month = 01"))
	fmt.Println(t.Format("day = 02"))
	fmt.Println(t.Format("year month day = 2006-01-02"))
	fmt.Println(t.Format("hour = 15"))
	fmt.Println(t.Format("minute = 04"))
	fmt.Println(t.Format("second = 05"))

	//休眠
	i := 0
	for {
		i++
		fmt.Println("i = ", i)
		time.Sleep(time.Millisecond * 100) //100毫秒
		if i == 10 {
			break
		}
	}
	fmt.Printf("Unix时间戳 = %v UnixNano = %v\n", t.Unix(), t.UnixNano())
	start := time.Now().Unix()
	exercise()
	end := time.Now().Unix()
	fmt.Println("运行时间 = ", end-start)
}

//练习 获取函数运行时间
func exercise() {
	str := " "
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
