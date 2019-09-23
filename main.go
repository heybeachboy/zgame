package main

import "zgame/service"

func main() {
	 game := service.NewGame()
	 game.PrintMartrix()
	 game.PushBlackChessmen(0,3)
	 game.PushBlackChessmen(3,2)
	 game.PushWhiteChessmen(2,3)
	 game.PushWhiteChessmen(3,5)
	 game.PrintMartrix()

}
