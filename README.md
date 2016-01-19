# wargame
In this game players take turns capturing empty squares until the board is filled. Each square has a point value. If the square captured is adjacent to an enemy square that is flanked (that is, surrounded by another ally piece) that square is captured. This is checked for all adjacent squares. After all squares have been filled the game ends and the player with the most points wins

#### To run the game

Use go run main.go -h to run the game manually

Use -v when running manually to get each move printed to console

./playgame.sh runs game.sh for every board in src/game_boards

./game.sh BOARD MMDEPTH ABDEPTH - Runs the all variations of the game with given board and depths.
