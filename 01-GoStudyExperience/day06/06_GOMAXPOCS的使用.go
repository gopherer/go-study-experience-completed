package main

import (
	"fmt"
	"runtime"
)

func main06() {
	// n := runtime.GOMAXPROCS(1)//指定以1线程---核运算
	n := runtime.GOMAXPROCS(4) //指定以4核运算
	fmt.Println(n)

	for {
		go fmt.Print(1)
		fmt.Print(0)
	}

}

//CPU的核心数是指物理上，也就是硬件上存在着几个核心,
//比如，双核就是包括2个相对独立的CPU核心单元组，四核就包含4个相对独立的CPU核心单元组，等等，依次类推。
//线程数是一种逻辑的概念，简单地说，就是模拟出的CPU核心数。
//比如，可以通过一个CPU核心数模拟出2线程的CPU，也就是说，这个单核心的CPU被模拟成了一个类似双核心CPU的功能。我们从任务管理器的性能标签页中看到的是两个CPU。
