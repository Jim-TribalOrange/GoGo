package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	G := game{nextTurn: 'b'}
	G.gameBoard = board{size: 13}

	G.createBoard()

}
