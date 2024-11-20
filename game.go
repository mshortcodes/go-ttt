package main

import "fmt"

type gameBoard struct {
	moves [9]rune
	p1    string
	p2    string
	turn  string
}

func (b gameBoard) printMoves() {
	for i, move := range b.moves {
		switch i {
		case 2, 5, 8:
			fmt.Printf("%c\n", move)
		default:
			fmt.Printf("%c | ", move)
		}
	}
}

func (b *gameBoard) printTurn() {
	fmt.Printf("%s's turn > ", b.turn)
	if b.turn == b.p1 {
		b.turn = b.p2
		return
	}
	b.turn = b.p1
}
