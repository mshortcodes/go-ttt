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
	fmt.Println()
	for i, move := range b.moves {
		switch i {
		case 2, 5, 8:
			fmt.Printf("%c\n", move)
		default:
			fmt.Printf("%c | ", move)
		}
	}
	fmt.Println()
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
	nextMoveInt, err := strconv.Atoi(nextMoveStr)
	if err != nil {
		return 0, fmt.Errorf("couldn't convert to integer")
	}
	nextMoveIdx := nextMoveInt - 1
	return nextMoveIdx, nil
}

func (b *board) checkCell(idx int) error {
	if idx > len(b.moves)-1 || idx < 0 {
		return fmt.Errorf("cell is out of bounds")
	}

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
