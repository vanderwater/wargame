package wargame

import (
	"math"
)

/* Using Wikipedia's psuedocode for minimax as an outline

function minimax(node, depth, maximizingPlayer)
    if depth = 0 or node is a terminal node
        return the heuristic value of node
    if maximizingPlayer
        bestValue := -∞
        for each child of node
            val := minimax(child, depth - 1, FALSE)
            bestValue := max(bestValue, val)
        return bestValue
    else
        bestValue := +∞
        for each child of node
            val := minimax(child, depth - 1, TRUE)
            bestValue := min(bestValue, val)
        return bestValue
*/

func minimaxMove(board *Board, depth int, currplayer string, minplayer string, maxplayer string) (int, int, int) {
	var bestValue int
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
					val, _, _ := minimaxMove(board, depth-1, minplayer, minplayer, maxplayer)
					if val > bestValue || bestValue == 0 {
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
					val, _, _ := minimaxMove(newBoard, depth-1, maxplayer, minplayer, maxplayer)
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
