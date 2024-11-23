package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getNextMove(scanner *bufio.Scanner) int {

}

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
