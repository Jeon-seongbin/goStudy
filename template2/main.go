package main

import "fmt"

func PlayGame(
	start func(),
	takeTurn func(),
	haveWinner func() bool,
	winningPlayer func() int) {

	start()

	for !haveWinner() {
		takeTurn()
	}

	fmt.Printf("Player %d wins \n", winningPlayer())

}

func main() {
	turn := 1
	maxTurns := 10
	currentPlayer := 0

	start := func() {
		fmt.Println("Starting a game of chess")
	}

	takeTurn := func() {
		turn++
		fmt.Printf("Turn %d taken by player %d \n", turn, currentPlayer)
		currentPlayer = (currentPlayer + 1) % 2
	}

	haveWinner := func() bool {
		return turn == maxTurns
	}

	winningPlayer := func() int {
		return currentPlayer
	}

	PlayGame(start, takeTurn, haveWinner, winningPlayer)
}
