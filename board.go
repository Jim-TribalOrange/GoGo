package main

type board struct {
	size       int
	positions  []rune
	koPosition int
	moveError  string
}

func (b *board) addPiece(p *position, c rune) (bool, int) {

	//check if its a valid move

	// if already occupied - false
	po := getArrayPositionInt(p, b.size)

	if b.positions[po] != '0' {
		b.moveError = "Position taken"
		return false, 0
	}

	isKo := false

	// not Ko restriction
	if b.koPosition == po {
		b.moveError = "Ko"
		return false, 0
	}

	//is Ko
	if isKo {
		b.koPosition = po
	}

	b.positions[po] = c

	return true, 0
}
