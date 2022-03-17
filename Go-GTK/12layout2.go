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
	if _, err := builder.AddFromFile("12test.glade"); err != nil{
		fmt.Println(err)
	}
	//获取windows控件
	win := gtk.WindowFromObject(builder.GetObject("windows"))
	win.Connect("destroy",gtk.MainQuit)
	//获取布局
	layout := gtk.TableFromObject(builder.GetObject("table1"))
	//新建button
	button := gtk.NewButtonWithLabel("hello layout")
	layout.AttachDefaults(button,0,1,1,3)
	//显示控件
	win.ShowAll()
	//主事件循环
	gtk.Main()
}
