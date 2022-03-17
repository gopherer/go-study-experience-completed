package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"strconv"
	"unsafe"
)

//枚举，标志黑子白子状态
const (
	Empty = iota //当前棋盘格子没有子
	Black        //当前棋盘格子为黑子
	White        //当前棋盘格子白白
)

//控件结构体
type ChessWidget struct {
	window      *gtk.Window //窗口
	buttonMin   *gtk.Button //最小化窗口按钮
	buttonClose *gtk.Button //关闭窗口按钮
	labelBlack  *gtk.Label  //记录黑棋个数
	labelWhite  *gtk.Label  //记录白棋个数
	labelTime   *gtk.Label  //记录倒计时
	imageBlack  *gtk.Image  //提示黑子落子
	imageWhite  *gtk.Image  //提示白子落子
}

//控件属性结构体
type ChessInfo struct {
	w, h           int //窗口的宽度和高度
	x, y           int //鼠标点击，相当于窗口的坐标
	startX, startY int //棋盘起点坐标
	gridW, gridH   int //棋盘格子的宽度和高度
}

//黑白棋结构体
type ChessBoard struct {
	ChessWidget //匿名字段
	ChessInfo
	currentRole    int       //该谁落子
	tipTimerId     int       //实现提示闪烁效果的定时器id
	leftTimerId    int       //倒计时定时器id
	machineTimerID int       //机器落子定时器
	timeNum        int       //记录剩余时间
	chess          [8][8]int //二维数组，标记棋盘状态
}

func main() {
	//初始化
	gtk.Init(&os.Args)

	var obj ChessBoard //创建结构体变量

	obj.CreateWindow() //创建控件，设置控件属性
	obj.HandleSignal() //事件，信号处理
	obj.InitChess()    //初始化棋盘

	obj.window.Show() //显示控件

	//主事件循环
	gtk.Main()
}

//方法：创建控件，设置控件属性
func (obj *ChessBoard) CreateWindow() {
	//加载glade文件
	builder := gtk.NewBuilder()
	if _, err := builder.AddFromFile("UI.glade"); err != nil {
	}
	//窗口相关
	obj.window = gtk.WindowFromObject(builder.GetObject("window1")) //获取窗口
	obj.window.SetAppPaintable(true)                                //允许绘图
	obj.window.SetPosition(gtk.WIN_POS_CENTER)                      //显示居中
	obj.window.GetSize(&obj.w, &obj.h)                              //得到窗口宽度和高度
	obj.window.SetDecorated(false)                                  //去边框
	obj.window.SetIconFromFile("../image/icon.jpg")                 //设置游戏图标
	//设置事件，让窗口可以捕获鼠标点击和移动
	obj.window.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))
	//按钮相关
	obj.buttonMin = gtk.ButtonFromObject(builder.GetObject("buttonMin"))
	obj.buttonClose = gtk.ButtonFromObject(builder.GetObject("buttonClose"))
	//给按钮设置图片
	ButtonSetImageFromFile(obj.buttonMin, "../image/min.png")
	ButtonSetImageFromFile(obj.buttonClose, "../image/close.png")
	//标签相关
	obj.labelBlack = gtk.LabelFromObject(builder.GetObject("labelBlack"))
	obj.labelWhite = gtk.LabelFromObject(builder.GetObject("labelWhite"))
	obj.labelTime = gtk.LabelFromObject(builder.GetObject("labelTime"))
	//设置字体大小
	obj.labelBlack.ModifyFontSize(70)
	obj.labelWhite.ModifyFontSize(70)
	obj.labelTime.ModifyFontSize(45)
	//设置内容
	obj.labelBlack.SetText("2")
	obj.labelWhite.SetText("2")
	obj.labelTime.SetText("20")
	//改变字体颜色
	obj.labelBlack.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("blue"))
	obj.labelWhite.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("red"))
	obj.labelTime.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white"))
	//image相关
	obj.imageBlack = gtk.ImageFromObject(builder.GetObject("imageBlack"))
	obj.imageWhite = gtk.ImageFromObject(builder.GetObject("imageWhite"))
	//设置图片
	ImageSetPicFromFile(obj.imageBlack, "../image/black.png")
	ImageSetPicFromFile(obj.imageWhite, "../image/white.png")
	//棋盘相关
	obj.startX, obj.startY = 410, 120
	obj.gridW, obj.gridH = 100, 80
}

//方法：事件，信号处理
func (obj *ChessBoard) HandleSignal() {
	//鼠标点击事件
	//"button-press-event" :鼠标按下是触发
	obj.window.Connect("button-press-event", MousePressEvent, obj)
	//鼠标移动事件
	//"motion-notify-event"：鼠标按下触发
	obj.window.Connect("motion-notify-event", MouseMoveEvent, obj)
	//按钮的信号处理
	obj.buttonMin.Clicked(func() {
		obj.window.Iconify() //最小化窗口
	})
	obj.buttonClose.Clicked(func() {
		glib.TimeoutRemove(obj.tipTimerId) //关闭定时器
		glib.TimeoutRemove(obj.leftTimerId)
		gtk.MainQuit() //关闭窗口
	})
	//绘图相关
	//大小改变事件
	//"configure_event"窗口大小改变时触发
	obj.window.Connect("configure_event", func() {
		obj.window.QueueDraw() //重新刷图
	})
	//绘图事件 “expose-event"
	obj.window.Connect("expose-event", PaintEvent, obj)
}

//方法:初始化棋盘，黑白棋辅助属性相关
func (obj *ChessBoard) InitChess() {
	//初始化棋盘
	for i := 0; i < len(obj.chess); i++ {
		for j := 0; j < len(obj.chess[0]); j++ {
			obj.chess[i][j] = Empty
		}
	}
	//默认有两个黑子和白子
	obj.chess[3][3] = Black
	obj.chess[4][4] = Black
	obj.chess[4][3] = White
	obj.chess[3][4] = White
	//重新游戏时，更新一下棋盘
	obj.window.QueueDraw()
	obj.labelBlack.SetText("2")
	obj.labelWhite.SetText("2")
	//image先隐藏
	obj.imageBlack.Hide()
	obj.imageWhite.Hide()
	//默认黑子先下
	obj.currentRole = Black
	//启动定时器
	obj.tipTimerId = glib.TimeoutAdd(500, func() bool {
		obj.ShowTip()
		return true
	})
	//倒计时定时器
	obj.timeNum = 20
	obj.labelTime.SetText(strconv.Itoa(obj.timeNum))
	obj.leftTimerId = glib.TimeoutAdd(1000, func() bool {
		obj.timeNum--
		obj.labelTime.SetText(strconv.Itoa(obj.timeNum))
		if obj.timeNum == 0 {
			obj.ChangeRole() //时间到，改变角色
		}
		return true
	})

}

//方法:提示功能,实现闪烁效果
func (obj *ChessBoard) ShowTip() {
	if obj.currentRole == Black { //当前黑子下
		//隐藏白子image
		obj.imageWhite.Hide()
		if obj.imageBlack.GetVisible() == true {
			//原来黑子是显示的，现在需要隐藏
			obj.imageBlack.Hide()
		} else {
			//原来黑子是隐藏的，现在需要显示
			obj.imageBlack.Show()
		}
	} else { //当前白子下
		//隐藏黑子image
		obj.imageBlack.Hide()
		if obj.imageWhite.GetVisible() == true {
			//原来白子是显示的，现在需要隐藏
			obj.imageWhite.Hide()
		} else {
			//原来白子是隐藏的，现在需要显示
			obj.imageWhite.Show()
		}
	}
}

//方法：改变落子角色
func (obj *ChessBoard) ChangeRole() {
	//改变角色重新设置时间
	obj.timeNum = 20
	obj.labelTime.SetText(strconv.Itoa(obj.timeNum))
	//先隐藏提示
	obj.imageBlack.Hide()
	obj.imageWhite.Hide()
	//原来三黑子下,就到白子下
	if obj.currentRole == Black {
		obj.currentRole = White
	} else {
		obj.currentRole = Black
	}
	obj.JudgeResult() //统计个数，判断胜负

	if obj.currentRole == White { //机器落子
		obj.machineTimerID = glib.TimeoutAdd(1500, func() bool {
			obj.MachinePlay() //机器落子
			return true
		})
	}
}

//方法:机器落子
func (obj *ChessBoard) MachinePlay() {
	//移除定时器
	glib.TimeoutRemove(obj.machineTimerID)

	max, px, py := 0, -1, -1

	//优先落子在四个角落,如果四个角落不能落子，就找吃子最多的位置
	for i := 0; i < len(obj.chess); i++ {
		for j := 0; j < len(obj.chess[0]); j++ {
			//最后一个参数为false，只判断能否吃子，不改变二维数组
			num := obj.JudgeRule(i, j, obj.currentRole, false)
			if num > 0 {
				if (i == 0 && j == 0) || (i == 7 && j == 0) || (i == 0 && j == 7) || (i == 7 && j == 7) {
					px, py = i, j
					goto End
				}
			}
			if num > max {
				max, px, py = num, i, j
			}
		}
	}
End:
	if px == -1 { //说明机器不能落子
		obj.ChangeRole() //改变角色
		return
	}
	//机器吃子
	obj.JudgeRule(px, py, obj.currentRole, true)
	obj.window.QueueDraw() //刷新整个窗口
	obj.ChangeRole()       //改变角色

}

//方法：吃子规则
//吃子规则的参数：棋盘数组坐标位置（x,y) role 当前落子角色
//eatChess为true，代表改变原来数组，false不改变数组内容，只判断此位置能吃多少个子
//返回值，吃子个数
func (obj *ChessBoard) JudgeRule(x, y int, role int, eatChess bool) (eatNum int) {
	// 棋盘的八个方向
	dir := [8][2]int{{1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	tempX, tempY := x, y                  // 临时保存棋盘数组坐标位置
	if obj.chess[tempX][tempY] != Empty { // 如果此方格内已有棋子，返回
		return 0
	}
	// 棋盘的8个方向
	for i := 0; i < 8; i++ {
		tempX += dir[i][0]
		tempY += dir[i][1] // 准备判断相邻棋子
		//如果没有出界，且相邻棋子为对方棋子，才有吃子的可能
		if (tempX < 8 && tempX >= 0 && tempY < 8 && tempY >= 0) && (obj.chess[tempX][tempY] != role) && (obj.chess[tempX][tempY] != Empty) {
			tempX += dir[i][0]
			tempY += dir[i][1] //继续判断下一个，向前走一步
			for tempX < 8 && tempX >= 0 && tempY < 8 && tempY >= 0 {
				if obj.chess[tempX][tempY] == Empty { //遇到空位跳出
					break
				}
				if obj.chess[tempX][tempY] == role { // 找到自己的棋子，代表可以吃子
					if eatChess == true { // 确定吃子
						obj.chess[x][y] = role // 开始点标志为自己的棋子
						tempX -= dir[i][0]
						tempY -= dir[i][1] // 后退一步
						for (tempX != x) || (tempY != y) {
							// 只要没有回到开始的位置就执行
							obj.chess[tempX][tempY] = role // 标志为自己的棋子
							tempX -= dir[i][0]
							tempY -= dir[i][1] // 继续后退一步
							eatNum++           // 累计
						}
					} else { //不吃子，只是判断这个位置能不能吃子
						tempX -= dir[i][0]
						tempY -= dir[i][1]                 // 后退一步
						for (tempX != x) || (tempY != y) { // 只计算可以吃子的个数
							tempX -= dir[i][0]
							tempY -= dir[i][1] // 继续后退一步
							eatNum++
						}
					}
					break // 跳出循环
				} // 没有找到自己的棋子，就向前走一步
				tempX += dir[i][0]
				tempY += dir[i][1] // 向前走一步
			}
		} // 如果这个方向不能吃子，就换一个方向
		tempX, tempY = x, y
	}
	return // 返回能吃子的个数
}

//方法：统计黑白棋个数，判断胜负
func (obj *ChessBoard) JudgeResult() {
	//判断胜负条件:双方都不能吃子
	result := ""
	isOver := true //默认游戏可以结束
	blackNum, whiteNum := 0, 0
	for i := 0; i < len(obj.chess); i++ {
		for j := 0; j < len(obj.chess[0]); j++ {
			if obj.chess[i][j] == Black {
				blackNum++
			} else if obj.chess[i][j] == White {
				whiteNum++
			}
			//只要一个位置能吃子，游戏还没有结束
			if obj.JudgeRule(i, j, Black, false) > 0 || obj.JudgeRule(i, j, White, false) > 0 {
				isOver = false
			}
		}
	}
	//界面显示个数
	obj.labelBlack.SetText(strconv.Itoa(blackNum))
	obj.labelWhite.SetText(strconv.Itoa(whiteNum))

	if isOver == false {
		return
	}
	//执行到这一步，说明游戏结束
	//关闭定时器
	glib.TimeoutRemove(obj.tipTimerId)
	glib.TimeoutRemove(obj.leftTimerId)

	//判断胜负
	if blackNum > whiteNum {
		result = "(我)蓝蜘蛛获得胜利\n是否继续游戏"
	} else if blackNum < whiteNum {
		result = "(机器)橙蜘蛛获得胜利\n是否继续游戏"
	} else {
		result = "平局\n是否继续游戏"
	}
	//问题对话框
	dialog := gtk.NewMessageDialog(
		obj.window,           //父窗口
		gtk.DIALOG_MODAL,     //模态对话框
		gtk.MESSAGE_QUESTION, //问题对话框
		gtk.BUTTONS_YES_NO,   //按钮
		result)               //对话框内容
	ret := dialog.Run()
	if ret == gtk.RESPONSE_YES { //Yes 继续游戏
		obj.InitChess() //重新初始化
	} else {
		gtk.MainQuit() //关闭窗口（游戏)
	}
	dialog.Destroy()
}

//函数：鼠标点击事件
func MousePressEvent(ctx *glib.CallbackContext) {
	//获取用户传递的参数
	obj, ok := ctx.Data().(*ChessBoard)
	if ok == false {
		fmt.Println("MousePressEvent ChessBoard err")
		return
	}
	//获取鼠标按下属性结构体，系统内部的变量，不是用户传参变量
	arg := ctx.Args(0)
	event := *(**gdk.EventButton)(unsafe.Pointer(&arg))
	//保存点击的x,y坐标
	obj.x, obj.y = int(event.X), int(event.Y)
	//fmt.Printf("相对于窗口的坐标 x = %d y = %d \n", obj.x, obj.y)
	//获取棋盘坐标
	i := (obj.x - obj.startX) / obj.gridW
	j := (obj.y - obj.startY) / obj.gridH
	if obj.currentRole == White { //如果是白棋下，用户不能点击
		return
	}
	if i >= 0 && i < len(obj.chess) && j >= 0 && j < len(obj.chess[0]) {
		//fmt.Println(i, j)
		//obj.chess[i][j] = obj.currentRole
		//吃子，落子必须能吃子
		if obj.JudgeRule(i, j, obj.currentRole, true) > 0 {
			obj.window.QueueDraw() //刷新整个窗口
			obj.ChangeRole()       //改变角色
		}

	}
}

//函数：鼠标移动事件
func MouseMoveEvent(ctx *glib.CallbackContext) {
	//获取用户传递的参数
	obj, ok := ctx.Data().(*ChessBoard)
	if ok == false {
		fmt.Println("MouseMoveEvent ChessBoard err")
		return
	}
	//获取鼠标按下属性结构体，系统内部的变量，不是用户传参变量
	arg := ctx.Args(0)
	event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

	x, y := int(event.XRoot)-obj.x, int(event.YRoot)-obj.y
	obj.window.Move(x, y) //移动窗口
}

//函数：给按钮设置图标
func ButtonSetImageFromFile(button *gtk.Button, fileName string) {
	//获取按钮的大小
	w, h := 0, 0
	button.GetSizeRequest(&w, &h)
	//创建pixBuf
	pixBuf, _ := gdkpixbuf.NewPixbufFromFileAtScale(fileName, w, h, false)
	//创建image
	image := gtk.NewImageFromPixbuf(pixBuf)
	//释放pixBuf
	pixBuf.Unref()
	//给按钮设置图片
	button.SetImage(image)
	//去掉按钮的焦距
	button.SetCanFocus(false)
}

//函数:绘图事件
func PaintEvent(ctx *glib.CallbackContext) {
	//获取用户传递的参数
	obj, ok := ctx.Data().(*ChessBoard)
	if ok == false {
		fmt.Println("PaintEvent ChessBoard err")
		return
	}
	//获取画家
	painter := obj.window.GetWindow().GetDrawable()
	//设置绘图区域
	gc := gdk.NewGC(painter)
	//新建背景PixBuf
	bgPixBuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("../image/bg.jpg", obj.w, obj.h, false)
	//黑白子PixBuf
	blackPixBuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("../image/black.png", obj.gridW, obj.gridH, false)
	whitePixBuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("../image/white.png", obj.gridW, obj.gridH, false)
	//绘图
	painter.DrawPixbuf(gc, bgPixBuf, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
	//画黑白棋
	for i := 0; i < len(obj.chess); i++ {
		for j := 0; j < len(obj.chess[0]); j++ {
			if obj.chess[i][j] == Black {
				painter.DrawPixbuf(gc, blackPixBuf, 0, 0, obj.startX+i*obj.gridW, obj.startY+j*obj.gridH, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
			} else if obj.chess[i][j] == White {
				painter.DrawPixbuf(gc, whitePixBuf, 0, 0, obj.startX+i*obj.gridW, obj.startY+j*obj.gridH, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
			}
		}
	}
	//释放资源
	bgPixBuf.Unref()
}

//函数：给image设置图片
func ImageSetPicFromFile(image *gtk.Image, filename string) {
	//获取image大小
	w, h := 0, 0
	image.GetSizeRequest(&w, &h)
	//创建pixBuf
	pixBuf, _ := gdkpixbuf.NewPixbufFromFileAtScale(filename, w, h, false)
	//给image设置图片
	image.SetFromPixbuf(pixBuf)
	//释放pixBuf
	pixBuf.Unref()
}
