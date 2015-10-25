package wargame

import (
	"fmt"
)

func switchPlayers(currPlayer, minPlayer, maxPlayer string) string {
	if currPlayer == maxPlayer {
		return minPlayer
	} else {
		return maxPlayer
	}
}

func PlayGame(board *Board) {
	var maxPlayer string = "Green"
	var minPlayer string = "Blue"
	currPlayer := maxPlayer

	for !board.isGameOver() {
		_, x, y := minimaxMove(board, 3, currPlayer, minPlayer, maxPlayer)
		if x == -1 && y == -1 {
			fmt.Print("Error\n")
		}
		board.capturePoint(x, y, currPlayer)
		fmt.Printf("%v: %v, %v\n", currPlayer, x, y)
		board.Print()

		currPlayer = switchPlayers(currPlayer, minPlayer, maxPlayer)
	}
	fmt.Print(board.CalculateAllScores())
}
