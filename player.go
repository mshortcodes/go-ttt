package main

import (
	"bufio"
	"fmt"
)

type player struct {
	name   string
	symbol rune
}

func createPlayer(scanner *bufio.Scanner, playerNum int) player {
	fmt.Printf("Enter name for player %d: ", playerNum)
	scanner.Scan()
	name := scanner.Text()
	return player{
		name:   name,
		symbol: 'X',
	}
}
