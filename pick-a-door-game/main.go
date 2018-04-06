package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type game struct {
	Rounds int
}

func main() {
	game := game{
		Rounds: 100,
	}
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for input != "y" && input != "n" {
		fmt.Print("Would you like to change the number of rounds calculated? (Type 'y' or 'n')\n")
		scanner.Scan()
		input = scanner.Text()
		if input != "y" && input != "n" {
			fmt.Print("Please follow the directions...")
		}
	}

	if input == "y" {
		var integer int
		var err error
		for integer == 0 {
			fmt.Print("How many rounds would you like to run?\n")
			scanner.Scan()
			input = scanner.Text()
			integer, err = strconv.Atoi(input)
			if integer == 0 {
				fmt.Print("Please enter an integer that isn't 0...\n")
			} else if err != nil {
				fmt.Print("Please enter an integer...\n")
			}
		}
		game.setGameRounds(integer)
	}

	currentRounds := 0
	for currentRounds <= game.Rounds {
		currentRounds++
	}
}

func (g *game) setGameRounds(rounds int) {
	g.Rounds = rounds
}
