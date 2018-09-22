package main

type group struct {
	positions   []int
	liabilities []int
}

func (b *board) createGroup(p int, colour rune) group {

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

func (g *group) liabilityHeld(p int) bool {

	for _, heldPos := range g.liabilities {

		if heldPos == p {
			return true
		}
	}
	return false
}
