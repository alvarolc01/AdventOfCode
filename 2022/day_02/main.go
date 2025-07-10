package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Move int
type Outcome int

const (
	Rock Move = iota
	Paper
	Scissors
)

const (
	Lose Outcome = iota
	Tie
	Win
)

var moveScores = []int{1, 2, 3}
var outcomeScores = []int{0, 3, 6}

type RPSGame struct {
	opponentMove Move
	myMove       Move
}

func charToMove(char byte) (Move, error) {
	switch char {
	case 'A', 'X':
		return Rock, nil
	case 'B', 'Y':
		return Paper, nil
	case 'C', 'Z':
		return Scissors, nil
	default:
		return 0, fmt.Errorf("invalid char: %c", char)
	}
}

func NewRPSGame(line string) (*RPSGame, error) {
	if len(line) < 3 {
		return nil, fmt.Errorf("invalid game: %s", line)
	}
	opponentMove, err := charToMove(line[0])
	if err != nil {
		return nil, err
	}

	mymove, err := charToMove(line[2])
	if err != nil {
		return nil, err
	}

	return &RPSGame{
		opponentMove: opponentMove,
		myMove:       mymove,
	}, nil
}

func result(myMove, oppMove Move) Outcome {
	if myMove == oppMove {
		return Tie
	}
	if (myMove-oppMove+3)%3 == 1 {
		return Win
	}
	return Lose
}

func moveForOutcome(oppMove Move, outcome Outcome) Move {
	if outcome == Tie {
		return oppMove
	} else if outcome == Win {
		return (oppMove + 1) % 3
	} else {
		return (oppMove + 2) % 3
	}
}

func part1(games []RPSGame) {
	total := 0

	for _, game := range games {
		outcome := result(game.myMove, game.opponentMove)
		total += outcomeScores[outcome] + moveScores[game.myMove]
	}

	fmt.Printf("Part 1: %d\n", total)
}

func part2(games []RPSGame) {
	total := 0
	for _, game := range games {
		var requiredOutcome Outcome
		if game.myMove == Rock {
			requiredOutcome = Lose
		} else if game.myMove == Paper {
			requiredOutcome = Tie
		} else if game.myMove == Scissors {
			requiredOutcome = Win
		}

		requiredMove := moveForOutcome(game.opponentMove, requiredOutcome)
		total += outcomeScores[requiredOutcome] + moveScores[requiredMove]
	}

	fmt.Printf("Part 2: %d\n", total)

}

func parseInput(lines []string) ([]RPSGame, error) {
	games := make([]RPSGame, len(lines))
	for idx, line := range lines {
		game, err := NewRPSGame(line)
		if err != nil {
			return nil, err
		}

		games[idx] = *game
	}

	return games, nil
}

func main() {
	fileName := flag.String("file", "", "Path to the file to read")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("file name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	fileContent, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("error reading input file:", err)
		os.Exit(1)
	}
	lines := strings.Split(string(fileContent), "\n")

	games, err := parseInput(lines)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(games)
	part2(games)
}
