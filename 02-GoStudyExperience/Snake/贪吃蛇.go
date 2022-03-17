package main

import (
	"02_gocode/Snake/Clib"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//全局常量 界面大小
const WIDE int = 20
const HIGH int = 20

//定义全局变量食物
var food Food

//定义全局分数
var score int = 0

//初始化父类 坐标
type Position struct {
	X int
	Y int
}

//初始化子类 蛇
type Snake struct {
	size int                   //长度
	dir  byte                  //方向
	pos  [WIDE * HIGH]Position //定义数组存储每一节蛇的坐标
}

//初始化子类 食物
type Food struct {
	Position
}

//初始化蛇信息
func (s *Snake) SnakeInit() {
	s.size = 2
	//方向 R 右 L左 U 上 D 下
	s.dir = 'R'
	//定义蛇头
	s.pos[0].X = WIDE / 2
	s.pos[0].Y = HIGH / 2
	//定义蛇身
	s.pos[1].X = WIDE/2 - 1
	s.pos[1].Y = HIGH / 2
	//显示蛇的位置
	for i := 0; i < s.size; i++ {
		if i == 0 {
			ShowUI(s.pos[i].X, s.pos[i].Y, '@')
		} else {
			ShowUI(s.pos[i].X, s.pos[i].Y, '*')
		}
	}
	//go 添加一个独立的函数 非阻塞执行
	//接收键盘中的输入
	go func() {
		for {
			switch Clib.Direction() {
			//根据键盘输入设置方向
			//上
			case 72, 87, 119:
				if s.dir != 'D' {
					s.dir = 'U'
				}
			//下
			case 80, 83, 115:
				if s.dir != 'U' {
					s.dir = 'D'
				}
			//左
			case 65, 75, 97:
				if s.dir != 'R' {
					s.dir = 'L'
				}
			//右
			case 68, 77, 100:
				if s.dir != 'L' {
					s.dir = 'R'
				}
			//空格 暂停
			case 32:
				s.dir = 'P'
			}
		}
	}()
}

//游戏开始
func (s *Snake) PlayGame() {
	//更新蛇坐标的偏移值
	dx, dy := 0, 0
	//游戏流程控制
	for {
	FLAG:
		//延迟执行
		time.Sleep(time.Second / 3)
		if s.dir == 'P' {
			goto FLAG
		}
		//更新蛇的位置
		switch s.dir {
		case 'U':
			dx, dy = 0, -1
		case 'D':
			dx, dy = 0, 1
		case 'L':
			dx, dy = -1, 0
		case 'R':
			dx, dy = 1, 0
		}
		//蛇和墙壁膨胀
		if s.pos[0].X <= 0 || s.pos[0].X >= WIDE || s.pos[0].Y < 0 || s.pos[0].Y > HIGH-3 {
			return
		}
		//蛇和身体碰撞
		for i := 1; i < s.size; i++ {
			if s.pos[0].X == s.pos[i].X && s.pos[0].Y == s.pos[i].Y {
				return
			}
		}
		//蛇身和食物重合 重置食物
		for i := 1; i < s.size; i++ {
			if food.X == s.pos[i].X && food.Y == s.pos[i].Y {
				RandomFood()
				i = 1
			}
		}
		//蛇头和食物碰撞
		if s.pos[0].X == food.X && s.pos[0].Y == food.Y {
			//身体增长
			s.size++
			//随机新食物
			RandomFood()
			score++
		}
		//记录末尾坐标
		lx := s.pos[s.size-1].X
		ly := s.pos[s.size-1].Y
		//更新蛇的坐标 蛇身体的坐标
		for i := s.size; i > 0; i-- {
			s.pos[i].X = s.pos[i-1].X
			s.pos[i].Y = s.pos[i-1].Y
		}
		//蛇头的坐标
		s.pos[0].X += dx
		s.pos[0].Y += dy
		//显示蛇的位置
		for i := 0; i < s.size; i++ {
			if i == 0 {
				ShowUI(s.pos[i].X, s.pos[i].Y, '@')
			} else {
				ShowUI(s.pos[i].X, s.pos[i].Y, '*')
			}
		}
		//将末尾置为空
		ShowUI(lx, ly, ' ')
	}
}

//随机食物
func RandomFood() {
	//随机食物
	food.X = rand.Intn(WIDE) //0-19
	food.Y = rand.Intn(HIGH)
	//显示食物位置
	ShowUI(food.X, food.Y, '#')
}

//初始化地图   ？？？？？？
func MapInit() {
	//输出初始画面
	fmt.Fprintf(os.Stderr, `

  &---------------------------------------&
  |                                       |
  |                                       |
  |                                       |
  |                                       |
  |                                       |
  |                                       |
  |                                       |
  |                                       |
  |                                       |
  |                                       |
  |                                       |
  |                                       |
  |                                       |
  |                                       |
  |                                       |
  |                                       |
  |                                       |
  |                                       |
  &---------------------------------------&

`)
}

//展示界面
func ShowUI(X int, Y int, ch byte) {
	//调用C语言代码 设置控制台光标
	Clib.GotPosition(X*2+2, Y+3)
	//将字符绘制在界面中
	fmt.Fprintf(os.Stderr, "%c", ch)
}

func main() {
	//设置随机数种子 用于混淆
	rand.Seed(time.Now().UnixNano())
	//隐藏控制台光标
	Clib.HideCursor()
	//初始化地图
	MapInit()
	//随机食物
	RandomFood()
	//创建蛇对象
	var s Snake
	s.SnakeInit()
	s.PlayGame()
	//打印分数
	Clib.GotPosition(0, 22)
	fmt.Fprintln(os.Stderr, "score = ", score)
	Clib.Pause()
}
