package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	//gtk初始化
	gtk.Init(&os.Args)
	//加载glade文件
	builder := gtk.NewBuilder()
	if _, err := builder.AddFromFile("04test.glade"); err !=nil{
		fmt.Println(err)
	}
	//获得glade上的控件
	win := gtk.WindowFromObject(builder.GetObject("windows"))
	b1 := gtk.ButtonFromObject(builder.GetObject("button"))
	b1.Clicked(func() {
		fmt.Println("button1 已被按下")
	})
	win.Connect("destroy", func() {
		fmt.Println("程序已经关闭")
		gtk.MainQuit()
	})
	//显示控件，如果通过glade添加的控件，show即可显示所以
	//        如果通过代码布局添加的控件，需要ShowAll才能显示所以
	win.Show()

	//主事件循环
	gtk.Main()
}

