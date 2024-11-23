package main

import (
	"bufio"
	"fmt"
	"strconv"
)

type board struct {
	moves      [9]rune
	p1         player
	p2         player
	nextPlayer string
	scanner    *bufio.Scanner
}

func (b board) printBoard() {
	for i, move := range b.moves {
		switch i {
		case 2, 5, 8:
			fmt.Printf("%c\n", move)
		default:
			fmt.Printf("%c | ", move)
		}
	}
}

func (b board) printTurn() {
	fmt.Printf("%s's turn > ", b.nextPlayer)
}

func (b *board) toggleTurn() {
	switch {
	case b.nextPlayer == b.p1.name:
		b.nextPlayer = b.p2.name
	case b.nextPlayer == b.p2.name:
		b.nextPlayer = b.p1.name
	}
}

func (b *board) getNextMove() (int, error) {
	b.scanner.Scan()
	nextMoveStr := b.scanner.Text()
	nextMoveIdx, err := strconv.Atoi(nextMoveStr)
	if err != nil {
		return 0, fmt.Errorf("couldn't convert to integer")
	}
	return nextMoveIdx, nil
}

func (b *board) checkCell(idx int) error {
	if b.moves[idx] != 0 {
		return fmt.Errorf("invalid move")
	}
	return nil
}
func (b *board) addMove(idx int) {
	switch {
	case b.nextPlayer == b.p1.name:
		b.moves[idx] = b.p1.symbol
	case b.nextPlayer == b.p2.name:
		b.moves[idx] = b.p2.symbol
	}
}
