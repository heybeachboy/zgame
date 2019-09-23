package service

import (
	"fmt"
	"time"
	"golang.org/x/exp/rand"
	)

const (
	EMPTY     = iota //表示棋盘位置为空
	WHITE            //表示棋盘位置为白棋
	BLACK            //表示棋盘位置为黑棋
	PLAYERNUM = 2    //玩家数量
	ROW       = 6    //棋盘的行数
	COL       = 7    //棋盘的列数
)

func init() {
	rand.Seed(uint64(time.Now().Unix()))
}

/**
 *创建一个游戏
 */
func NewGame() (*Game) {
	g := Game{}
	g.Init()
	return &g
}

type Game struct {
	Player        map[int]int //玩家数量如果玩家的个人信息加载，可以用LIST或者SLICE的数据结构
	currentPlayer int         //当前玩家
	Martrix       [][]int     //保存棋盘信息
	row           int         //棋盘行数
	col           int         //棋盘列数
	count         int         //统计棋盘使用了的数量
	total         int         //棋盘总共的位置
}

func (g *Game) Init() { //该位置都是用常量配置，为了简便就不做传参处理
	g.initPlayer()
	g.initMartrix(ROW, COL)
	g.currentPlayer = 1
	g.row = ROW
	g.col = COL
	g.count = 0
	g.total = ROW * COL
}

/**
 *初始化玩家信息
 */
func (g *Game) initPlayer() {
	g.Player = make(map[int]int, PLAYERNUM)
	g.Player[1] = BLACK
	g.Player[2] = WHITE

}

/**
 *初始化棋盘
 */
func (g *Game) initMartrix(row, col int) {
	for i := 0; i < row; i++ {
		g.Martrix = append(g.Martrix, make([]int, COL))
	}
}

/**
 *检查棋盘位置是不是合法
 */
func (g *Game) isIllegal(row, col int) (bool) {

	if row < 0 || col < 0 || row >= g.row || col >= g.col {
		return false
	}
	return true
}

/**
 *判断棋子的类型是不是合法，
 */
func (g *Game) isChessmenOfType(chessmen int) (bool) {

	if chessmen < 0 || chessmen > 2 {
		return false
	}

	return true

}

/**
 *往棋盘放入棋子
 */
func (g *Game) PushChessmen(row, col, chessmen int) {

	if !g.isChessmenOfType(chessmen) || !g.isIllegal(row, col) {
		panic(fmt.Sprintf("传入的位置或者棋子不合法 row : %d, col : %d, chessman : %d", row, col, chessmen))
	}

	g.Martrix[row][col] = chessmen
	g.count++

}

/**
 *往棋盘对应位置放入黑棋
 */
func (g *Game) PushBlackChessmen(row, col int) {
	g.PushChessmen(row, col, BLACK)
}

/**
 *往棋盘位置放入白棋
 */

func (g *Game) PushWhiteChessmen(row, col int) {
	g.PushChessmen(row, col, WHITE)

}

/*
 *打印输出棋盘信息
 */
func (g *Game) PrintMartrix() {

	for i := 0; i < g.row; i++ {
		for j := 0; j < g.col; j++ {
			fmt.Printf("%d ", g.Martrix[i][j])
		}
		fmt.Print("\n")
	}

	fmt.Printf("total : %d count : %d \n", g.total, g.count)

}

/**
 *
 */
func (g *Game) StartRandomFillMartrixGame() {
	g.currentPlayer = 1
	for {
		r := g.getRandomRow()
		if !g.judgeRowIsEmpty(r) { //判断该列是不是还存在空位
			g.swapPlayer() //如果不存在交换玩家再进行下一次寻找
			continue
		}
		g.filMartrix(r)
		g.swapPlayer()

		if g.total == g.count {
			break
		}

	}
}

/**
 *获取一个随机列
 */
func (g *Game) getRandomRow() (int) {
	return rand.Intn(g.row)

}

/**
 *判断该列是否存在空的位置
 */
func (g *Game) judgeRowIsEmpty(row int) (bool) {
	return g.Martrix[row][g.col-1] == 0
}

/**
 *填充martrix
 */
func (g *Game) filMartrix(row int) {
	for k, v := range g.Martrix[row] {
		if v != 0 {
			continue
		}
		g.PushChessmen(row, k, g.Player[g.currentPlayer])
		break

	}
}

/**
 *交换出棋玩家
 */
func (g *Game) swapPlayer() {
	if g.currentPlayer == 1 {
		g.currentPlayer = 2
		return
	}
	g.currentPlayer = 1

}
