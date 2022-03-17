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
	if _, err := builder.AddFromFile("07test.glade"); err != nil {
		fmt.Println(err)
	}
	//获取控件
	win := gtk.WindowFromObject(builder.GetObject("windows"))
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	//加载图片控件
	image := gtk.ImageFromObject(builder.GetObject("image1"))
	//获取控件大小
	var w, h int
	image.GetSizeRequest(&w,&h)
	//设置一张图片资源，pixBuf 控件大小要和图片大小一样
	//false：不保存图片原来的尺寸
	pixBuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("png/y1.png",w,h,false)
	//给image设置图片
	image.SetFromPixbuf(pixBuf)
	//图片资源使用完毕，需要释放空间（底层C语言的要求）
	pixBuf.Unref()
	//显示控件
	win.Show()
	//主事件循环
	gtk.Main()
}
