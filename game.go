package main

import "fmt"

type gameBoard struct {
	moves      [9]rune
	p1         player
	p2         player
	nextPlayer string
}

type player struct {
	name   string
	symbol rune
}

func (b gameBoard) printBoard() {
	for i, move := range b.moves {
		switch i {
		case 2, 5, 8:
			fmt.Printf("%c\n", move)
		default:
			fmt.Printf("%c | ", move)
		}
	}
}

func (b gameBoard) printTurn() {
	fmt.Printf("%s's turn > ", b.nextPlayer)
}

func (b *gameBoard) toggleTurn() {
	switch {
	case b.nextPlayer == b.p1.name:
		b.nextPlayer = b.p2.name
	case b.nextPlayer == b.p2.name:
		b.nextPlayer = b.p1.name
	}
}

func (b *gameBoard) addMove(idx int) {
	switch {
	case b.nextPlayer == b.p1.name:
		b.moves[idx] = b.p1.symbol
	case b.nextPlayer == b.p2.name:
		b.moves[idx] = b.p2.symbol
	}
}
