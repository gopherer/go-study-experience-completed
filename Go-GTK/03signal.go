package main

import (
	"fmt"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	//初始化
	gtk.Init(&os.Args)
	//创建主窗口
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	win.SetSizeRequest(1000,1000)
	//创建容器控件（固定布局，任意布局）
	layout := gtk.NewFixed()
	//将布局添加到窗口
	win.Add(layout)
	//创建按钮
	b1 := gtk.NewButtonWithLabel("$_$")
	b2 := gtk.NewButtonWithLabel("*_*")
	b2.SetSizeRequest(303,330)
	//将按钮添加到布局中
	layout.Put(b1,0,0)
	layout.Put(b2,300,300)
	//显示控件
	win.ShowAll()
	//信号处理
	//按下按钮触发“clicked"
	str := "hello gtk"
	b1.Connect("clicked",HandleSignal,str)//也可以使用匿名函数推荐使用
	b2.Clicked(func() {
		fmt.Println(str+" go")
	})
	//按窗口关闭按钮，触发“destroy"
	win.Connect("destroy", func() {
		fmt.Println("触发destroy")
		gtk.MainQuit()
	})
	//主事件循环
	gtk.Main()
	fmt.Println("程序已关闭")
}
func HandleSignal(ctx *glib.CallbackContext){
	fmt.Println("hello ")
	tep := ctx.Data()//空接口，接收用户传递的数据
	data,ok := tep.(string)
	if ok {
		fmt.Println(data)
	}
}
