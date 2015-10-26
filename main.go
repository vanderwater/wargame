package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/vanderwater/wargame/src"
	"os"
	"strings"
)

//TODO/vanderwater: Error Check these flags after Parse
var (
	// These are all pointers
	board          = flag.String("board", "", "Board to select")
	maxStratString = flag.String("p1strat", "mm", "Choose \"mm\" or \"ab\"")
	minStratString = flag.String("p2strat", "mm", "Choose \"mm\" or \"ab\"")
	maxDepth       = flag.Int("p1depth", 3, "Choose a number 1-5")
	minDepth       = flag.Int("p2depth", 3, "Choose a number 1-5")
)

func determineStrategy(strat string) func(*wargame.Board, int, string, string, string) (int, int, int) {
	if strings.ToLower(strat) == "mm" || strings.ToLower(strat) == "minimax" {
		return wargame.MinimaxMove
	} else {
		return wargame.AlphabetaMove
	}
}

func main() {

	flag.Parse()

	if *board != "" {
		*board = strings.Title(strings.ToLower(*board))
		boardFile, err := os.Open("src/game_boards/" + *board + ".txt")
		if err == nil {

			maxStrat := determineStrategy(*maxStratString)
			minStrat := determineStrategy(*minStratString)
			board := wargame.ReadBoard(bufio.NewReader(boardFile))
			board.Print()
			wargame.PlayGame(board, maxStrat, minStrat, *maxDepth, *minDepth)
			return
		}
	}

	fmt.Print("Choose from boards: Sevastopol, Keren, Narvik, Smolensk, Westerplatte\n")
}
