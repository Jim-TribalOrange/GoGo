package main

type position struct {
	row, col, pos    int
	colour           rune
	surrounding      []position
	connections      []int
	ownliabilities   []int
	oppositionStones []int
	status           int
}

func (p *position) getSurrounding(b *board, colour rune) {

	boardSize := b.size
	p.surrounding = make([]position, 4)
	po := getArrayPositionInt(p, boardSize)

	up := po - boardSize
	down := po + boardSize
	left := po - 1
	right := po + 1

	//Up
	if po-boardSize > -1 {
		p.surrounding[0] = positionFromInt(up, boardSize)
		p.surrounding[0].status = b.checkPosition(up, colour)
	} else {
		p.surrounding[0].status = 5
	}

	//down
	if po+boardSize < (boardSize * boardSize) {
		p.surrounding[1] = positionFromInt(down, boardSize)
		p.surrounding[1].status = b.checkPosition(down, colour)
	} else {
		p.surrounding[1].status = 5
	}
	//left
	if po%(boardSize+1) > 0 {
		p.surrounding[2] = positionFromInt(left, boardSize)
		p.surrounding[2].status = b.checkPosition(left, colour)
	} else {
		p.surrounding[2].status = 5
	}
	//Right - modus position with size, if < 19 or modus 19 then on right edge of the board
	if po < boardSize && po%boardSize > 0 {
		p.surrounding[3] = positionFromInt(right, boardSize)
		p.surrounding[3].status = b.checkPosition(right, colour)
	} else {
		p.surrounding[3].status = 5
	}

	for _, sp := range p.surrounding {
		switch sp.status {
		case 0:
			p.ownliabilities = append(p.ownliabilities, sp.pos)
		case 1:
			p.connections = append(p.connections, sp.pos)
		case 2:
			p.oppositionStones = append(p.oppositionStones, sp.pos)
		}

	}

}

func (b *board) checkPosition(check int, c rune) int {

	switch b.positions[check] {
	case 0:
		return 0
	case c:
		return 1
	default:
		return 2
	}

}

func getArrayPositionInt(p *position, size int) int {

	res := (size * (p.row - 1)) + p.col

	//make it array-able
	return res - 1
}

func positionFromInt(po int, boardsize int) position {

	po++

	y := po / boardsize
	x := po % boardsize
	return position{row: x, col: y, pos: po}

}
