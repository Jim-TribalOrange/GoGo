package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	G := game{turn: 'b'}
	G.board = board{size: 13}

	G.createBoard()

}

type position struct {
	row, col    int
	colour      rune
	connections []position
	liabilities int
}

type board struct {
	size       int
	positions  []position
	koPosition int
	moveError  string
}

type game struct {
	board board
	turn  rune
}

func (g *game) createBoard() {

	s := g.board.size

	g.board.positions = make([]position, s*s)
	positionX := 1
	positionY := 1

	for i := range g.board.positions {

		if positionY > s {
			positionX++
			positionY = 1
		}

		if positionX > s {
			break
		}

		g.board.positions[i] = position{row: positionX, col: positionY}

		positionY++
	}

	fmt.Printf("%+v\n", g.board)
}

func (b *board) addPiece(p position, c rune) (bool, int) {

	//check if its a valid move

	// if already occupied - false
	po := getPositionInt(p, b.size)

	if b.positions[po].colour != 0 {
		b.moveError = "Position taken"
		return false, 0
	}

	isKo := false

	// not Ko restriction
	if b.koPosition == po {
		b.moveError = "Ko"
		return false, 0
	}
	connections, liabilities, others := b.checkMove(po, c)

	//is Ko
	if isKo {
		b.koPosition = po
	}

	b.positions[po].colour = c

	return true, 0
}

func (b *board) checkMove(po int, colour rune) ([]position, []position, []position) {

	//check row + 1 and row - 1 & col + 1 and col -1
	//need to recurse neighbours of the same colour (dont double count)
	connections := make([]position, 4)
	liabilities := make([]position, 4)
	other := make([]position, 4)
	check := make([]int, 4)
	up := po - b.size
	down := po + b.size
	left := po - 1
	right := po + 1

	//Up
	if po-b.size > -1 {
		check[0] = b.checkPosition(up, colour)
	} else {
		check[0] = 5
	}

	//down
	if po+b.size < (b.size * b.size) {
		check[1] = b.checkPosition(down, colour)
	} else {
		check[1] = 5
	}
	//left
	if po%(b.size+1) > 0 {
		check[2] = b.checkPosition(left, colour)
	} else {
		check[2] = 5
	}
	//Right - modus position with size, if < 19 or modus 19 then on right edge of the board
	if po < b.size && po%b.size > 0 {
		check[3] = b.checkPosition(right, colour)
	} else {
		check[3] = 5
	}

	for i := range check {
		switch i {
		case -1:

		case 0:
			liabilities = append(b.positions[po-b.size])
		case 1:
			fmt.Println("Linux.")
		case 2:
		default:

		}

		return connections, liabilities, other
	}
}

func (b *board) checkPosition(chk int, c rune) int {

	var rtn int
	switch b.positions[chk].colour {
	case 0:
		rtn = 0
	case c:
		rtn = 1
	default:
		rtn = 2
	}

	return rtn
}

func getPositionInt(p position, size int) int {
	//Size(X-1) + Y
	res := ((p.col - 1) * size) + p.row

	//make it array-able
	return res - 1
}

func getPosition(position int, size int) (int, int) {

	position += 1

	y := position / size
	x := position % size

	return y, x
}

func (g *game) checkCaptures(p position, colour rune) (bool, int) {

	//if there is a capture it will involve a neigbouring position

	//check there are liabilities first - only need to worry if there arent any
	liabilities := g.getLiabilities(p)

	return false, 0

}
