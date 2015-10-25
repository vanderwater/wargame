package main

import (
	"bufio"
	"flag"
	"fmt"
	. "github.com/vanderwater/wargame/src"
	"os"
)

var (
	// These are all pointers
	board = flag.String("board", "", "Board to select")
)

func main() {

	flag.Parse()

	if *board == "" {
		fmt.Print("Choose from boards: Sevastopol, Keren, Narvik, Smolensk, Westerplatte\n")
	} else {
		boardFile, _ := os.Open("src/game_boards/" + *board + ".txt")
		board := ReadBoard(bufio.NewReader(boardFile))
		board.Print()
		PlayGame(board)
	}
}
