package service

import "fmt"

const (
	EMPTY     = iota //表示棋盘位置为空
	WHITE            //表示棋盘位置为白棋
	BLACK            //表示棋盘位置为黑棋
	PLAYERNUM = 2    //玩家数量
	ROW       = 6    //棋盘的行数
	COL       = 7    //棋盘的列数
	)

/**
 *创建一个游戏
 */
func NewGame() (*Game) {
	g := Game{}
	g.Init()
	return &g
}

type Game struct {
	Player  int     //玩家数量如果玩家的个人信息加载，可以用LIST或者SLICE的数据结构
	Martrix [][]int //保存棋盘信息
	row     int     //棋盘行数
	col     int     //棋盘列数
	count   int     //统计棋盘使用了的数量
	total   int     //棋盘总共的位置
}

func (g *Game) Init() { //该位置都是用常量配置，为了简便就不做传参处理
	g.Player = PLAYERNUM
	g.initMartrix(ROW, COL)
	g.row = ROW
	g.col = COL
	g.count = 0
	g.total = ROW * COL
}

/**
 *初始化棋盘
 */
func (g *Game) initMartrix(row, col int) {
	for i := 0; i < row; i++ {
		g.Martrix = append(g.Martrix,make([]int, COL))
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
		panic(fmt.Sprintf("传入的位置或者棋子不合法 row : %d, col : %d, chessman : %d",row,col,chessmen))
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

	fmt.Print("\n")

}
