package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter name for player 1: ")
	scanner.Scan()
	p1 := scanner.Text()

	fmt.Println("Enter name for player 2: ")
	scanner.Scan()
	p2 := scanner.Text()

	board := gameBoard{
		p1:   p1,
		p2:   p2,
		turn: p1,
	}

	for {
		board.printMoves()
		board.printTurn()
	}
}
