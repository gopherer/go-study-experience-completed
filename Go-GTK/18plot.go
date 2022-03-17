package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	//初始化
	gtk.Init(&os.Args)
	//加载glade文件
	builder := gtk.NewBuilder()
	if _, err := builder.AddFromFile("18test.glade"); err != nil {
		fmt.Println(err)
	}
	x := 0
	//获取控件
	win := gtk.WindowFromObject(builder.GetObject("windows"))
	win.Connect("destroy", gtk.MainQuit)
	//"configure-event"窗口改变是触发
	var w, h int
	win.Connect("configure-event", func() {
		win.GetSize(&w, &h)
		//刷图，整个窗口区域刷图
		win.QueueDraw() //触发绘图信号，“expose-event"
	})
	//允许窗口绘图
	win.SetAppPaintable(true)
	//绘图时触发的信号，"expose-event"
	win.Connect("expose-event", func() {
		//设置画家
		painter := win.GetWindow().GetDrawable()
		//指定绘图区域
		gc := gdk.NewGC(painter)
		//创建图片资源
		bg, _ := gdkpixbuf.NewPixbufFromFileAtScale("png/photo.jpg", w, h, false)
		flower, _ := gdkpixbuf.NewPixbufFromFileAtScale("png/y1.png", 80, 50, false)

		painter.DrawPixbuf(gc, bg, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
		painter.DrawPixbuf(gc, flower, 0, 0, x, 150, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
		//释放图片资源
		bg.Unref()
		flower.Unref()
	})
	//获取按钮
	button := gtk.ButtonFromObject(builder.GetObject("button1"))
	button.Clicked(func() {
		x += 50
		if x > w {
			x = 0
		}
		//刷新绘图
		win.QueueDraw()
	})
	//显示控件
	win.Show()
	//主事件循环
	gtk.Main()
}
