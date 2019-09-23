package service
const (
	  PLAYERNUM = 2 //玩家数量
	  ROW = 6 //棋盘的行数
	  COL = 7 //棋盘的列数
	  EMPTY = iota //表示棋盘位置为空
	  WHITE        //表示棋盘位置为白棋
	  BLACK        //表示棋盘位置为黑棋
)

type Game struct {
	 Player int
	 Martrix [][]int
	 row int
	 col int
	 count int
	 Total int
}

func (g *Game)Init() { //该位置都是用常量配置，为了简便就不做传参处理
     g.Player = PLAYERNUM
     g.initMartrix(ROW,COL)
     g.row = ROW
     g.col = COL
     g.count = 0
     g.Total = ROW * COL
     }

/**
 *初始化棋盘
 */
func (g *Game)initMartrix(row, col int) {
	 for i :=0; i < row; i++ {
	 	 g.Martrix[i] = make([]int,COL)
	 }
}

func (g *Game)isOverStep(row,col int)(bool) {
	 if row >= g.row || col >= g.col {
	 	return false
	 }
	 return true
}

func (g *Game)PushChessmen(row,col,chessmen int) {

}


func (g *Game)PushBlackChessmen(row,col int) {

}

func (g *Game)PushWhiteChessmen(row,col int) {

}

func (g *Game)PrintMartrix() {

}


