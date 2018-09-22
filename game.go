package main

type game struct {
	gameBoard     board
	nextTurn      rune
	turns         []turn
	blackCaptures int
	whiteCaputres int
}

func (g *game) createBoard() {

	s := g.gameBoard.size

	g.gameBoard.positions = make([]rune, s*s)

}

func (g *game) checkCaptures(p position, colour rune) (bool, int) {

	//if there is a capture it will involve a neigbouring position

	//check there are liabilities first - only need to worry if there arent any
	//liabilities := g.getLiabilities(p)

	return false, 0

}
