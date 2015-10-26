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

// TODO/vanderwater: Pass strategies as functions, use helper functions to limit arguments
func PlayGame(board *Board, maxPlayerStrategy, minPlayerStrategy func(*Board, int, string, string, string) (int, int, int), maxPlayerDepth, minPlayerDepth int) {
	var maxPlayer string = "Green"
	var minPlayer string = "Blue"
	currPlayer := maxPlayer

	for !board.isGameOver() {
		var x, y int
		// Get Move
		if currPlayer == maxPlayer {
			_, x, y = maxPlayerStrategy(board, maxPlayerDepth, currPlayer, minPlayer, maxPlayer)
		} else {
			_, x, y = minPlayerStrategy(board, minPlayerDepth, currPlayer, minPlayer, maxPlayer)
		}

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
