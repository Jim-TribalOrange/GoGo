package main

type board struct {
	size          int
	positions     []rune
	koPosition    int
	blackCaptures int
	whiteCaputres int
	moveError     string
}

func (b *board) addPiece(p int, c rune) (bool, int) {

	//check if its a valid move

	// if already occupied - false

	if b.positions[p] != '0' {
		b.moveError = "Position taken"
		return false, 0
	}

	// not Ko restriction
	if b.koPosition == p {
		b.moveError = "Ko"
		return false, 0
	}

	isCapture, newCaptures := b.checkCaputres(p, c)

	if isCapture {
		switch c {
		case ('b'):
			b.blackCaptures += newCaptures
		case ('w'):
			b.whiteCaputres += newCaptures
		}
	}

	po := positionFromInt(p, b.size)
	po.getSurrounding(b, c)

	if !isCapture && len(po.ownliabilities) == 0 {
		b.moveError = "Suicide"
		return false, 0
	}

	b.positions[p] = c
	return true, 0
}

func (b *board) checkCaputres(p int, c rune) (bool, int) {

	pos := positionFromInt(p, b.size)

	pos.getSurrounding(b, c)
	capture := false
	newCaputres := 0
	for _, po := range pos.oppositionStones {

		g := b.createGroup(po, notC(c))

		if len(g.liabilities) == 0 {

			b.koPosition = p
			newCaputres = newCaputres + b.removeGroup(g)

			capture = true
		}
	}

	return capture, newCaputres

}

func (b *board) removeGroup(g group) int {

	counter := 0

	for _, p := range g.positions {

		b.positions[p] = 0
		counter++
	}

	return counter
}

func notC(c rune) rune {

	switch c {
	case 'b':
		return 'w'
	case 'w':
		return 'b'
	default:
		return 0
	}
}
