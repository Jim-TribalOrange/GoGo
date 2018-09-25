package main

type game struct {
	gameBoard board
	nextTurn  rune
	turns     []turn
}

func (g *game) createBoard() {

	s := g.gameBoard.size

	g.gameBoard.positions = make([]rune, s*s)

}

//check there are liabilities first - only need to worry if there arent any
//liabilities := g.getLiabilities(p)
