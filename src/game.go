package wargame

import (
	"fmt"
	"math"
	"time"
)

var (
	maxNodes, minNodes int
)

// Return the opposite player than what it currently is
func switchPlayers(currPlayer, minPlayer, maxPlayer string) string {
	if currPlayer == maxPlayer {
		return minPlayer
	} else {
		return maxPlayer
	}
}

// Plays the wargame. Prints info to console and returns nothing
// TODO/vanderwater: Do not hardcode maxPlayer and minPlayer names, replace them somehow
func PlayGame(board *Board, maxPlayerStrategy, minPlayerStrategy func(*Board, int, string, string, string) (int, int, int), maxPlayerDepth, minPlayerDepth int, verbose bool) {
	var maxPlayer string = "Blue"
	var minPlayer string = "Green"
	var maxTime, minTime time.Duration
	currPlayer := maxPlayer

	for !board.isGameOver() {
		var x, y int
		// Get move of corresponding player
		if currPlayer == maxPlayer {
			currentTime := time.Now()
			_, x, y = maxPlayerStrategy(board, maxPlayerDepth, currPlayer, minPlayer, maxPlayer)
			maxTime += time.Since(currentTime)
		} else {
			currentTime := time.Now()
			_, x, y = minPlayerStrategy(board, minPlayerDepth, currPlayer, minPlayer, maxPlayer)
			minTime += time.Since(currentTime)
		}

		if x == -1 && y == -1 {
			fmt.Print("Error: Depth must be greater than 0\n")
		}

		board.capturePoint(x, y, currPlayer)
		if verbose {
			fmt.Printf("%v: %v, %v\n", currPlayer, x, y)
			board.Print()
		}

		currPlayer = switchPlayers(currPlayer, minPlayer, maxPlayer)
	}
	// Print ending status of board
	board.Print()
	fmt.Print(board.CalculateAllScores())
	fmt.Printf("%v expanded %v nodes\n%v expanded %v nodes\n", maxPlayer, maxNodes, minPlayer, minNodes)
	fmt.Printf("%v spent %.4f seconds deciding\n%v spent %.4f seconds deciding\n", maxPlayer, maxTime.Seconds(), minPlayer, minTime.Seconds())
}

/** MINIMAX **/
/* Using Wikipedia's psuedocode for Minimax as an outline

function Minimax(node, depth, maximizingPlayer)
    if depth = 0 or node is a terminal node
        return the heuristic value of node
    if maximizingPlayer
        bestValue := -∞
        for each child of node
            val := Minimax(child, depth - 1, FALSE)
            bestValue := max(bestValue, val)
        return bestValue
    else
        bestValue := +∞
        for each child of node
            val := Minimax(child, depth - 1, TRUE)
            bestValue := min(bestValue, val)
        return bestValue
*/

func MinimaxMove(board *Board, depth int, currplayer string, minPlayer string, maxPlayer string) (int, int, int) {
	expandMax := currplayer == maxPlayer
	return MinimaxMoveAux(board, depth, currplayer, minPlayer, maxPlayer, expandMax)
}

func MinimaxMoveAux(board *Board, depth int, currplayer string, minPlayer string, maxPlayer string, expandMax bool) (int, int, int) {
	var bestValue int
	var bestx, besty int
	// Increment Expanded Nodes
	if expandMax {
		maxNodes++
	} else {
		minNodes++
	}

	// Heuristic is max - min, max / min, and max are also possibilities
	if depth == 0 || board.isGameOver() {
		return board.CalculatePlayerScore(maxPlayer) / (board.CalculatePlayerScore(minPlayer) + 1), -1, -1
	}
	if currplayer == maxPlayer {
		bestValue = math.MinInt32
		for y, row := range board.points {
			for x := range row {
				// If a piece can be played, test it
				if board.isValidPoint(x, y) && !board.isOccupiedPoint(x, y) {
					newBoard := board.Copy()
					newBoard.capturePoint(x, y, maxPlayer)
					val, _, _ := MinimaxMoveAux(newBoard, depth-1, minPlayer, minPlayer, maxPlayer, expandMax)
					if val > bestValue {
						bestValue = val
						bestx = x
						besty = y
					}
				}
			}
		}
		return bestValue, bestx, besty
	} else {
		bestValue = math.MaxInt32
		for y, row := range board.points {
			for x := range row {
				// If a piece can be played, test it
				if board.isValidPoint(x, y) && !board.isOccupiedPoint(x, y) {
					newBoard := board.Copy()
					newBoard.capturePoint(x, y, minPlayer)
					val, _, _ := MinimaxMoveAux(newBoard, depth-1, maxPlayer, minPlayer, maxPlayer, expandMax)
					if val < bestValue {
						bestValue = val
						bestx = x
						besty = y
					}
				}
			}
		}
		return bestValue, bestx, besty
	}
}

/** AlphaBeta Pruning **/
/**
function alphabeta(node, depth, α, β, maximizingPlayer)
      if depth = 0 or node is a terminal node
          return the heuristic value of node
      if maximizingPlayer
          v := -∞
          for each child of node
              v := max(v, alphabeta(child, depth - 1, α, β, FALSE))
              α := max(α, v)
              if β ≤ α
                  break (* β cut-off *)
          return v
      else
          v := ∞
          for each child of node
              v := min(v, alphabeta(child, depth - 1, α, β, TRUE))
              β := min(β, v)
              if β ≤ α
                  break (* α cut-off *)
          return v
*/

func AlphabetaMove(board *Board, depth int, currplayer, minPlayer, maxPlayer string) (int, int, int) {
	expandMax := currplayer == maxPlayer
	return alphabetaMoveAux(board, depth, math.MinInt32, math.MaxInt32, currplayer, minPlayer, maxPlayer, expandMax)
}

func alphabetaMoveAux(board *Board, depth int, alpha int, beta int, currplayer string, minPlayer string, maxPlayer string, expandMax bool) (int, int, int) {
	var bestValue int
	var bestx, besty int
	// Increment Expanded Nodes
	if expandMax {
		maxNodes++
	} else {
		minNodes++
	}

	if depth == 0 || board.isGameOver() {
		return board.CalculatePlayerScore(maxPlayer) / (board.CalculatePlayerScore(minPlayer) + 1), -1, -1
	}
	// Max player section
	if currplayer == maxPlayer {
		bestValue = math.MinInt32
	ExpandMax:
		for y, row := range board.points {
			for x := range row {
				// If a piece can be played, test it
				if board.isValidPoint(x, y) && !board.isOccupiedPoint(x, y) {
					newBoard := board.Copy()
					newBoard.capturePoint(x, y, maxPlayer)
					val, _, _ := alphabetaMoveAux(newBoard, depth-1, alpha, beta, minPlayer, minPlayer, maxPlayer, expandMax)
					if val > alpha {
						alpha = val
					}
					if val > bestValue {
						bestValue = val
						bestx = x
						besty = y
					}
					if beta <= alpha {
						break ExpandMax
					}
				}
			}
		}
		return bestValue, bestx, besty

	} else {
		//Min Player section
		bestValue = math.MaxInt32
	ExpandMin:
		for y, row := range board.points {
			for x := range row {
				// If a piece can be played, test it
				if board.isValidPoint(x, y) && !board.isOccupiedPoint(x, y) {
					newBoard := board.Copy()
					newBoard.capturePoint(x, y, minPlayer)
					val, _, _ := alphabetaMoveAux(newBoard, depth-1, alpha, beta, maxPlayer, minPlayer, maxPlayer, expandMax)
					if val < beta {
						beta = val
					}
					if val < bestValue {
						bestValue = val
						bestx = x
						besty = y
					}
					if beta <= alpha {
						break ExpandMin
					}
				}
			}
		}
		return bestValue, bestx, besty
	}
}
