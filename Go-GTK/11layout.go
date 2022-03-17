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
	if _, err := builder.AddFromFile("11test.glade"); err !=nil{
		fmt.Println(err)
	}
	//获取控件
	win := gtk.WindowFromObject(builder.GetObject("windows"))
	win.Connect("destroy",gtk.MainQuit)
	//获取布局
	layout := gtk.HBoxFromObject(builder.GetObject("hbox1"))
	//新建按钮
	button := gtk.NewButton()
	button.SetLabel("hello gtk")
	layout.Add(button)
	//显示控件
	win.ShowAll()
	//主事件循环
	gtk.Main()
}
