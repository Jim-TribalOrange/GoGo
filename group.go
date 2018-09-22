package main

type group struct {
	positions   []int
	liabilities []int
}

func (b *board) createGroup(p int, colour rune) group {

	//populate the group from the position listed
	grp := group{positions: make([]int, 4), liabilities: make([]int, 4)}
	posit := positionFromInt(p, b.size)
	posit.getSurrounding(b, colour)

	for _, aPos := range posit.connections {
		grp.positions = append(grp.positions, aPos)
		grp.getConnectedPosition(aPos, colour, b)
	}
	return grp
}

//need to recurse over positionsHeld, getting new positions

func (g *group) getConnectedPosition(p int, colour rune, b *board) {

	posit := positionFromInt(p, b.size)
	posit.getSurrounding(b, colour)

	added := make([]int, 4)

	for _, aPos := range posit.connections {
		if !g.positionHeld(aPos) {
			g.positions = append(g.positions, aPos)
			added = append(added, aPos)
		}
	}

	if len(added) > 0 {
		for _, aPos := range added {

			g.getConnectedPosition(aPos, colour, b)
		}

	}

}

func (g *group) positionHeld(p int) bool {

	for _, heldPos := range g.positions {

		if heldPos == p {
			return true
		}
	}
	return false
}
