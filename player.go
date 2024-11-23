package main

import (
	"bufio"
	"fmt"
)

type player struct {
	name   string
	symbol string
}

func createPlayer(playerNum int, symbol string, scanner *bufio.Scanner) player {
	name := ""

	for {
		fmt.Printf("Enter name for player %d: ", playerNum)
		scanner.Scan()
		name = scanner.Text()
		if name == "" {
			fmt.Println("Invalid name")
			continue
		}
		break
	}

	return player{
		name:   name,
		symbol: symbol,
	}
}
