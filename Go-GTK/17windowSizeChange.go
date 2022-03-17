package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	//初始化
	gtk.Init(&os.Args)
	//加载glade文件
	builder := gtk.NewBuilder()
	if _, err := builder.AddFromFile("17test.glade"); err != nil {
		fmt.Println(err)
	}
	//获取控件
	win := gtk.WindowFromObject(builder.GetObject("windows"))
	win.Connect("destroy", gtk.MainQuit)
	//"configure_event"窗口大小改变时触发
	win.Connect("configure_event", func() {
		var w, h int
		win.GetSize(&w, &h)
		fmt.Printf("w = %v, h = %v\n", w, h)
	})
	//显示控件
	win.Show()
	//主事件循环
	gtk.Main()
}
