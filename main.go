package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	G := game{size: 13}
	G.createBoard()

}

type position struct {
	row, col int
	colour   rune
}

type game struct {
	size       int
	board      []position
	turn       rune
	koPosition int
	moveError  string
}

func (g *game) createBoard() {

	s := g.size

	g.board = make([]position, s*s)
	positionX := 1
	positionY := 1

	for i := range g.board {

		if positionY > s {
			positionX++
			positionY = 1
		}

		if positionX > s {
			break
		}

		g.board[i] = position{row: positionX, col: positionY}

		positionY++
	}

	fmt.Printf("%+v\n", g.board)
}

func (g *game) addPiece(p position, c rune) (bool, int) {

	//check if its a valid move

	// if already occupied - false
	po := getposition(p)

	if g.board[po].colour != 0 {
		g.moveError = "Position taken"
		return false, false
	}

	isKo := false

	// not Ko restriction
	if g.koPosition == po {
		g.moveError = "Ko"
		return false, false
	}

	//is Ko
	if isKo {
		g.koPosition = p
	}

	// if no liabilities (after capture) - false
	//if there is a capture then there must be a liability
	ok, caps := checkCaptures(p)

	//if ok add piece return true
	g.board[po].colour = c

	return true, caps
}

func getLiabilities(position p, b *board) int {

	size = Sqrt(len(b))

	//check row + 1 and row - 1 & col + 1 and col -1
	//need to recurse neighbours of the same colour (dont double count)
	return 0
}

func getposition(p position, size int) int {

	//Size(X-1) + Y
	res := ((p.col - 1) * size) + p.row

	//make it array-able
	return res - 1

}

func checkCaptures(position p, colour rune) (bool, int) {

	//if there is a capture it will involve a neigbouring position

	//check there are liabilities first - only need to worry if there arent any
	liabilities := getLiabilities(p)

	return false, 0

}
