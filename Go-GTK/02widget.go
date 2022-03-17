package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	//初始化
	gtk.Init(&os.Args)
	//创建主窗口
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("go Gtk")
	win.SetSizeRequest(800,800)
	//创建容器控件（固定布局，任意布局）
	layout := gtk.NewFixed()
	//将布局添加到窗口
	win.Add(layout)
	//创建按钮
	b1 := gtk.NewButtonWithLabel("^_^")
	b2 := gtk.NewButtonWithLabel("$_$")
	b2.SetSizeRequest(200,200)
	//将按钮添加到布局中
	layout.Put(b1,0,0)
	layout.Put(b2,30,30)
	//显示控件
	//win.Show()
	//layout.Show()
	//b1.Show()
	//b2.Show()
	win.ShowAll()
	//主事件循环
	gtk.Main()
}
