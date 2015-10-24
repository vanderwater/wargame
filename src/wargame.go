package main

import (
	"bufio"
	"fmt"
)

type PointInfo struct {
	owner string
	value int
}

func ReadBoard(r bufio.Reader) [][]PointInfo {
	result = make([][]PointInfo, 1)
	var y int
	for line, err := r.ReadString(); err == nil; line, err := r.ReadString() {
		values := strings.Split(line, " ")
		result[y] = append(result, make([]PointInfo, 0))
		for value := range values {
			parsed_value = strconv.ParseInt(value, 0, 64)
			result[y] = append(result[y], PointInfo{owner: "X", value: int(parsed_value)})
		}
		y++
	}
	return result
}

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
func PrintBoard(board [][]PointInfo) {
	result = make([]string, 0)
	for y := 0; y < len(board)*4; y++ {
		result = append(result, strings.Repeat("#", len(board[y])*5))
		result = append(result, "") // Sets up strings to be printed
		result = append(result, "") // Sets up strings to be printed
		for x := 0; x < len(board[y]); x++ {
			result[y*4+1] += "# " + string(board[y][x].owner[0]) + " #"
			result[y*4+2] += "#" + strings.Repeat(" ", 3-getWidth(board[y][x].value)) + strconv.FormatInt(board[y][x].value, 10) + "#"
		}
		result = append(result, strings.Repeat("#", len(board[y])*5))
	}

	fmt.Printf("%v\n", strings.Join(result, "\n"))
}

func main() {
	boardFile = os.Open("game_boards/Keren.txt")
	board := ReadBoard(boardFile)
	PrintBoard(board)
}
