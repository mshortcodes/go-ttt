package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	p1 := createPlayer(1, "X", scanner)
	p2 := createPlayer(2, "O", scanner)

	board := board{
		moves:         [9]string{},
		p1:            p1,
		p2:            p2,
		currentPlayer: p1.name,
		scanner:       scanner,
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

		isWinner := board.checkForWinner()
		if isWinner {
			board.printBoard()
			board.printWinner()
			return
		}

		isDraw := board.checkforDraw()
		if isDraw {
			board.printBoard()
			board.printDraw()
			return
		}

		board.toggleTurn()
	}
}
