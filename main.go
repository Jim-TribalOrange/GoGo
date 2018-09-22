package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	G := game{nextTurn: 'b'}
	G.gameBoard = board{size: 13}

	G.createBoard()

}

// func (b *board) readPosition(p position, colour rune) bool {

// 	po := getPositionInt(p, b.size)
// 	result := true

// 	for i := range p.surrounding {
// 		switch i.status {
// 		case -1:

// 		case 0:
// 			b.positions[po-b.size].liabilities++
// 		case 1:
// 			b.positions[po-b.size].connections
// 		case 2:
// 		default:

// 		}

// 		return result
// 	}
// }
