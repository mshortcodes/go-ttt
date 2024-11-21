package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter name for player 1: ")
	scanner.Scan()
	name := scanner.Text()
	p1 := player{
		name:   name,
		symbol: 'X',
	}

	fmt.Println("Enter name for player 2: ")
	scanner.Scan()
	name = scanner.Text()
	p2 := player{
		name:   name,
		symbol: 'O',
	}

	board := gameBoard{
		moves:      [9]rune{},
		p1:         p1,
		p2:         p2,
		nextPlayer: p1.name,
	}

	for {
		board.printBoard()
		board.printTurn()
		scanner.Scan()
		nextMoveCell := scanner.Text()
		nextMoveIdx, err := strconv.Atoi(nextMoveCell)
		if err != nil {
			fmt.Println("couldn't convert to integer")
			continue
		}
		board.addMove(nextMoveIdx)
		board.toggleTurn()
	}
}
