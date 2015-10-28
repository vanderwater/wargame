echo -e \\n Running with MM Depth $2 and AB Depth $3
echo -e \\n$1 Minimax Minimax \\n
go run main.go -board=$1 -p1strat=mm -p1depth=$2 -p2strat=mm -p2depth=$2
echo -e \\n$1 Minimax Alpha-beta \\n
go run main.go -board=$1 -p1strat=mm -p1depth=$2 -p2strat=ab -p2depth=$3
echo -e \\n$1 Alpha-beta Minimax \\n
go run main.go -board=$1 -p1strat=ab -p1depth=$3 -p2strat=mm -p2depth=$2
echo -e \\n$1 Alpha-beta Alpha-beta \\n
go run main.go -board=$1 -p1strat=ab -p1depth=$3 -p2strat=ab -p2depth=$3
