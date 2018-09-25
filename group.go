package main

import "sort"

type group struct {
	positions   []int
	liabilities []int
	caputured   bool
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

	sort.Ints(g.positions)

}

func removeDuplicateGroups(groups []group) []group {

	final := make([]group, len(groups))
	positions := make([]int, len(groups))

	for _, g := range groups {

		if !contains(positions, g.positions[0]) {
			positions = append(positions, g.positions[0])
			final = append(final, g)
		}
	}

	return final
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
