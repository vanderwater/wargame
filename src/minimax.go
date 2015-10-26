package wargame

import (
	"math"
)

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

func MinimaxMove(board *Board, depth int, currplayer string, minplayer string, maxplayer string) (int, int, int) {
	var bestValue int = -1
	var bestx, besty int
	if depth == 0 || board.isGameOver() {
		return board.CalculatePlayerScore(maxplayer), -1, -1
	}
	if currplayer == maxplayer {
		for y, row := range board.points {
			for x := range row {
				if board.isValidPoint(x, y) && !board.isOccupiedPoint(x, y) {
					newBoard := board.Copy()
					newBoard.capturePoint(x, y, maxplayer)
					val, _, _ := MinimaxMove(board, depth-1, minplayer, minplayer, maxplayer)
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
				if board.isValidPoint(x, y) && !board.isOccupiedPoint(x, y) {
					newBoard := board.Copy()
					newBoard.capturePoint(x, y, minplayer)
					val, _, _ := MinimaxMove(newBoard, depth-1, maxplayer, minplayer, maxplayer)
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
