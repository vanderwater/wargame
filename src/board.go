package wargame

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type pointInfo struct {
	owner string
	value int
}

type Board struct {
	points [][]pointInfo
}

// Reads a board from the file
func ReadBoard(r *bufio.Reader) *Board {
	result := make([][]pointInfo, 0)
	var y int
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		values := strings.Split(line, "\t")
		result = append(result, make([]pointInfo, 0))
		for _, value := range values {
			parsed_value, err := strconv.ParseInt(strings.TrimSpace(value), 0, 64)
			if err != nil {
				fmt.Printf("%v: Error %v\n", value, err)
			}
			result[y] = append(result[y], pointInfo{owner: "X", value: int(parsed_value)})
		}
		y++
	}
	return &Board{result}
}

// Used for printing nicely
func getWidth(x int) int {
	if x == 0 {
		return 1
	}

	var result int
	for ; x > 0; x /= 10 {
		result++
	}
	return result
}

// TODO/vanderwater: Improve by using byte buffers or copy for string concat
// Prints the state of the board to the console
func (board *Board) Print() {
	if board == nil || len(board.points) == 0 {
		return
	}
	result := make([]string, 0)
	var y int
	for ; y < len(board.points); y++ {
		result = append(result, strings.Repeat("#", len(board.points[y])*4+1))
		result = append(result, "") // Sets up strings to be printed
		result = append(result, "") // Sets up strings to be printed
		for x := 0; x < len(board.points[y]); x++ {
			result[y*3+1] += "# " + string(board.points[y][x].owner[0]) + " "
			result[y*3+2] += "#" + strings.Repeat(" ", 3-getWidth(board.points[y][x].value)) + strconv.FormatInt(int64(board.points[y][x].value), 10)
		}
		result[y*3+1] = result[y*3+1] + "#"
		result[y*3+2] = result[y*3+2] + "#"
	}
	result = append(result, strings.Repeat("#", len(board.points[0])*4+1))

	fmt.Printf("%v\n", strings.Join(result, "\n"))
}

// A bunch of functions to tell us something about a point on the board
func (board *Board) isValidPoint(x, y int) bool {
	return !(board == nil || board.points == nil || y < 0 || x < 0 || y >= len(board.points) || x >= len(board.points[y]))
}

func (board *Board) isOccupiedPoint(x, y int) bool {
	return board.isValidPoint(x, y) && board.points[y][x].owner != "X"
}

func (board *Board) isAllyOccupiedPoint(x, y int, owner string) bool {
	return board.isValidPoint(x, y) && board.points[y][x].owner == owner
}

func (board *Board) isEnemyOccupiedPoint(x, y int, owner string) bool {
	return board.isValidPoint(x, y) && board.points[y][x].owner != owner && board.points[y][x].owner != "X"
}

func (board *Board) isPointAssissted(x, y int, owner string) bool {
	return board.isValidPoint(x, y) && (board.isAllyOccupiedPoint(y+1, x, owner) || board.isAllyOccupiedPoint(y-1, x, owner) || board.isAllyOccupiedPoint(y, x+1, owner) || board.isAllyOccupiedPoint(y, x-1, owner))
}

// Captures a point on the board, propogate capture blitzes enemy pieces
func (board *Board) capturePoint(x, y int, owner string) bool {
	if !board.isValidPoint(x, y) || board.isOccupiedPoint(x, y) {
		return false
	}

	board.points[y][x].owner = owner
	if board.isPointAssissted(x, y, owner) {
		board.propogateCapture(x, y, owner)
	}
	return true
}

func (board *Board) propogateCapture(x, y int, owner string) {
	if board.isOccupiedPoint(x, y+1) {
		board.points[y+1][x].owner = owner
	}
	if board.isOccupiedPoint(x, y-1) {
		board.points[y-1][x].owner = owner
	}
	if board.isOccupiedPoint(x+1, y) {
		board.points[y][x+1].owner = owner
	}
	if board.isOccupiedPoint(x-1, y) {
		board.points[y][x-1].owner = owner
	}
}

// Finds score of possible captures
func (board *Board) propogateScorePotential(x, y int, owner string) int {
	if !board.isValidPoint(x, y) {
		return 0
	}
	var result int
	if board.isEnemyOccupiedPoint(x, y-1, owner) {
		result += board.points[y-1][x].value
	}
	if board.isEnemyOccupiedPoint(x, y-1, owner) {
		result += board.points[y-1][x].value
	}
	if board.isEnemyOccupiedPoint(x+1, y, owner) {
		result += board.points[y][x+1].value
	}
	if board.isEnemyOccupiedPoint(x-1, y, owner) {
		result += board.points[y][x-1].value
	}
	return result
}

// Finds potential score of move
func (board *Board) ScorePotential(x, y int, owner string) int {
	if !board.isValidPoint(x, y) || board.isOccupiedPoint(x, y) {
		return 0
	}

	var result int
	result += board.points[x][y].value
	if board.isPointAssissted(x, y, owner) {
		result += board.propogateScorePotential(x, y, owner)
	}
	return result
}

func (board *Board) CalculatePlayerScore(owner string) int {
	var result int
	for _, row := range board.points {
		for _, point := range row {
			if point.owner == owner {
				result += point.value
			}
		}
	}
	return result
}

// TODO/vanderwater: Move away from hardcoded calculations
func (board *Board) CalculateAllScores() string {
	/**	scores := make(map[string]int)
	for _, row := range board.points {
		for _, point := range row {
			scores[point.owner] += point.value
		}
	}

	var result string
	for owner, total := range scores {
		result += fmt.Sprintf("%v: %v\n", owner, total)
	}

	return result
	*/
	greenScore := board.CalculatePlayerScore("Green")
	blueScore := board.CalculatePlayerScore("Blue")
	return fmt.Sprintf("Blue: %v\nGreen: %v\n", blueScore, greenScore)
}

// Deep Copies the board
func (board *Board) Copy() *Board {
	newBoard := make([][]pointInfo, len(board.points))
	for index, row := range board.points {
		newBoard[index] = make([]pointInfo, len(row))
		for x, point := range row {
			newBoard[index][x] = pointInfo{owner: point.owner, value: point.value}
		}
	}
	return &Board{newBoard}
}

// Detects if any squares are left open
func (board *Board) isGameOver() bool {
	for _, row := range board.points {
		for _, point := range row {
			if point.owner == "X" {
				return false
			}
		}
	}
	return true
}
