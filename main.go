package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	p1 := createPlayer(scanner, 1)
	p2 := createPlayer(scanner, 2)

	board := board{
		moves:      [9]rune{},
		p1:         p1,
		p2:         p2,
		nextPlayer: p1.name,
		scanner:    scanner,
	}

	for {
		board.printBoard()
		board.printTurn()
		nextMoveIdx, err := board.getNextMove()
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = board.checkCell(nextMoveIdx)
		if err != nil {
			fmt.Println(err)
			continue
		}
		board.addMove(nextMoveIdx)
		board.toggleTurn()
	}
}
