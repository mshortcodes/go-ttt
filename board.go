package main

import (
	"bufio"
	"fmt"
	"strconv"
)

type board struct {
	moves         [9]string
	p1            player
	p2            player
	currentPlayer string
	scanner       *bufio.Scanner
}

func (b board) printBoard() {
	fmt.Println()
	for i, move := range b.moves {
		switch i {
		case 2, 5, 8:
			fmt.Printf("%s\n", move)
		default:
			fmt.Printf("%s | ", move)
		}
	}
	fmt.Println()
}

func (b board) printTurn() {
	fmt.Printf("%s's turn > ", b.currentPlayer)
}

func (b *board) toggleTurn() {
	switch {
	case b.currentPlayer == b.p1.name:
		b.currentPlayer = b.p2.name
	case b.currentPlayer == b.p2.name:
		b.currentPlayer = b.p1.name
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

	if b.moves[idx] != "" {
		return fmt.Errorf("invalid move")
	}
	return nil
}

func (b *board) addMove(idx int) {
	switch {
	case b.currentPlayer == b.p1.name:
		b.moves[idx] = b.p1.symbol
	case b.currentPlayer == b.p2.name:
		b.moves[idx] = b.p2.symbol
	}
}

func (b *board) checkForWinner() bool {
	possibleWins := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, win := range possibleWins {
		if b.moves[win[0]] != "" && b.moves[win[0]] == b.moves[win[1]] && b.moves[win[1]] == b.moves[win[2]] {
			return true
		}
	}
	return false
}

func (b *board) printWinner() {
	fmt.Printf("%s has won!\n", b.currentPlayer)
}

func (b *board) checkforDraw() bool {
	for _, move := range b.moves {
		if move == "" {
			return false
		}
	}
	return true
}

func (b board) printDraw() {
	fmt.Println("It's a draw!")
}
