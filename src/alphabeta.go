package wargame

import (
	"math"
)

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

func AlphabetaMove(board *Board, depth int, currplayer, minplayer, maxplayer string) (int, int, int) {
	return alphabetaMoveAux(board, depth, math.MinInt32, math.MaxInt32, currplayer, minplayer, maxplayer)
}

func alphabetaMoveAux(board *Board, depth int, alpha int, beta int, currplayer string, minplayer string, maxplayer string) (int, int, int) {
	var bestValue int = math.MinInt32
	var bestx, besty int
	if depth == 0 || board.isGameOver() {
		return board.CalculatePlayerScore(maxplayer), -1, -1
	}
	if currplayer == maxplayer {
	ExpandMax:
		for y, row := range board.points {
			for x := range row {
				if board.isValidPoint(x, y) && !board.isOccupiedPoint(x, y) {
					newBoard := board.Copy()
					newBoard.capturePoint(x, y, minplayer)
					val, _, _ := alphabetaMoveAux(newBoard, depth-1, alpha, beta, minplayer, minplayer, maxplayer)
					if val > alpha {
						alpha = val
					}
					if beta <= alpha {
						break ExpandMax
					}
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
	ExpandMin:
		for y, row := range board.points {
			for x := range row {
				if board.isValidPoint(x, y) && !board.isOccupiedPoint(x, y) {
					newBoard := board.Copy()
					newBoard.capturePoint(x, y, minplayer)
					val, _, _ := alphabetaMoveAux(newBoard, depth-1, alpha, beta, maxplayer, minplayer, maxplayer)
					if val < beta {
						beta = val
					}
					if beta <= alpha {
						break ExpandMin
					}
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
