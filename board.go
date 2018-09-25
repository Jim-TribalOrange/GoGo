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
	b.positions[p] = c
	pos.getSurrounding(b, c)
	capture := false
	newCaputres := 0
	caputredGroups := make([]group, 4)
	for _, po := range pos.oppositionStones {

		g := b.getGroup(po, notColour(c))

		if len(g.liabilities) == 0 {
			caputredGroups = append(caputredGroups, g)
		}
	}

	caputredGroups = removeDuplicateGroups(caputredGroups)

	for _, cgroups := range caputredGroups {

		capture = true
		newCaputres = newCaputres + b.removeGroup(cgroups)
	}
	b.positions[p] = 0
	return capture, newCaputres

}

func (b *board) getGroup(p int, colour rune) group {

	//populate the group from the position listed
	grp := group{positions: make([]int, 4), liabilities: make([]int, 4)}

	grp.positions = append(grp.positions, p)
	posit := positionFromInt(p, b.size)
	posit.getSurrounding(b, colour)

	for _, aPos := range posit.connections {
		grp.positions = append(grp.positions, aPos)
		grp.getConnectedPosition(aPos, colour, b)
	}

	for _, pos := range grp.positions {

		po := positionFromInt(pos, b.size)

		for _, lia := range po.ownliabilities {
			if !grp.liabilityHeld(lia) {
				grp.liabilities = append(grp.liabilities, lia)
			}
		}
	}

	return grp
}

func (b *board) removeGroup(g group) int {

	counter := 0

	for _, p := range g.positions {

		b.positions[p] = 0
		counter++
	}

	return counter
}
