package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	availableRedCubes   int = 12
	availableGreenCubes int = 13
	availableBlueCubes  int = 14
)

type GameTurn struct {
	Red   int
	Green int
	Blue  int
}

func NewGameTurn(turn string) (GameTurn, error) {
	var currentTurn GameTurn

	for rock := range strings.SplitSeq(turn, ",") {
		var colour string
		var count int

		_, err := fmt.Sscanf(strings.TrimSpace(rock), "%d %s", &count, &colour)
		if err != nil {
			return GameTurn{}, err
		}

		switch colour {
		case "red":
			currentTurn.Red = count
		case "blue":
			currentTurn.Blue = count
		case "green":
			currentTurn.Green = count
		}
	}
	return currentTurn, nil
}

type Game struct {
	Id    int
	Turns []GameTurn
}

func NewGame(line string) (*Game, error) {
	parts := strings.Split(line, ":")
	var gameID int
	_, err := fmt.Sscanf(parts[0], "Game %d", &gameID)
	if err != nil {
		return nil, err
	}

	turns := strings.Split(strings.TrimSpace(parts[1]), ";")
	turnsGame := make([]GameTurn, len(turns))
	for idxTurn, turn := range turns {
		currTurn, err := NewGameTurn(turn)
		if err != nil {
			return nil, err
		}

		turnsGame[idxTurn] = currTurn
	}
	return &Game{
		Id:    gameID,
		Turns: turnsGame,
	}, nil
}

func (g *Game) IsValid() bool {
	valid := true
	for _, turn := range g.Turns {
		if turn.Green > availableGreenCubes || turn.Blue > availableBlueCubes || turn.Red > availableRedCubes {
			valid = false
			break
		}
	}

	return valid
}

func part1(input []*Game) {
	sumCalibrationValues := 0

	for _, game := range input {
		valid := game.IsValid()
		if valid {
			sumCalibrationValues += game.Id
		}
	}

	fmt.Printf("Part 1: %d\n", sumCalibrationValues)
}

func (g *Game) GetMinStones() (minRed, minBlue, minGreen int) {
	for _, turn := range g.Turns {
		minBlue = max(minBlue, turn.Blue)
		minGreen = max(minGreen, turn.Green)
		minRed = max(minRed, turn.Red)
	}

	return
}

func part2(input []*Game) {
	sumCalibrationValues := 0

	for _, game := range input {
		minRed, minBlue, minGreen := game.GetMinStones()
		sumCalibrationValues += minRed * minBlue * minGreen
	}

	fmt.Printf("Part 2: %d\n", sumCalibrationValues)
}

func parseInput(input []string) ([]*Game, error) {
	games := make([]*Game, len(input))

	for idxLine, line := range input {
		currGame, err := NewGame(strings.TrimSpace(line))
		if err != nil {
			return nil, err
		}

		games[idxLine] = currGame
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
	input := strings.Split(string(fileContent), "\n")

	games, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}
	part1(games)
	part2(games)
}
