package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	//初始化
	gtk.Init(&os.Args)
	//加载glade文件
	builder := gtk.NewBuilder()
	if _, err := builder.AddFromFile("09test.glade"); err !=nil{
		fmt.Println(err)
	}
	//获取控件win
	win := gtk.WindowFromObject(builder.GetObject("windows"))
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	//获取按钮
	b1 := gtk.ButtonFromObject(builder.GetObject("button1"))
	b2 := gtk.ButtonFromObject(builder.GetObject("button2"))
	//设置按钮标签
	b1.SetLabel("^_^")
	//改变字体大小
	b1.SetLabelFontSize(13)
	//无法按下按钮
	b1.SetSensitive(false)
	//获取按钮b2的大小
	var w, h int
	b2.GetSizeRequest(&w,&h)
	//新建图片资源，大小要与b2大小相符
	pixBuf, _  := gdkpixbuf.NewPixbufFromFileAtScale("png/y2.png",w,h,false)
	//新建image
	image := gtk.NewImageFromPixbuf(pixBuf)
	//释放资源pixBuf
	pixBuf.Unref()
	//按钮设置图标
	b2.SetImage(image)
	//取消焦距
	b2.SetCanFocus(false)
	//显示控件
	win.Show()
	//主事件循环
	gtk.Main()
}
