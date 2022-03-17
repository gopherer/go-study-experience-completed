package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	//初始化
	gtk.Init(&os.Args)
	//加载3glade文件
	builder := gtk.NewBuilder()
	if _, err := builder.AddFromFile("04test.glade"); err != nil{
		fmt.Println(err)
	}
	//获取win控件
	win := gtk.WindowFromObject(builder.GetObject("windows"))
	win.SetTitle("Hello GTK") //设置标题
	win.SetIconFromFile("png/y1.png")//设在图标
	win.SetSizeRequest(480,320)//设在窗口最小大小
	win.SetPosition(gtk.WIN_POS_CENTER)//窗口居中显示
	//获取窗口大小
	var w, h int
	win.GetSize(&w,&h)
	fmt.Printf("w = %d, h = %d\n",w,h)
	//限制窗口缩放
	win.SetResizable(false)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	//显示控件
	win.Show()
	//主事件循环
	gtk.Main()
}
