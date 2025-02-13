package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const EmptyButton rune = '#'

type Position struct {
	xAxis  int
	yAxis  int
	keypad [][]rune
}

func (p *Position) Move(direction rune) {
	movements := map[rune][2]int{
		'L': {0, -1},
		'R': {0, 1},
		'U': {-1, 0},
		'D': {1, 0},
	}

	if movement, valid := movements[direction]; valid {
		nextX := p.xAxis + movement[0]
		nextY := p.yAxis + movement[1]

		xWithinBounds := nextX >= 0 && nextX < len(p.keypad)
		yWithinBounds := nextY >= 0 && nextY < len(p.keypad)

		if xWithinBounds && yWithinBounds && p.keypad[nextX][nextY] != EmptyButton {
			p.xAxis = nextX
			p.yAxis = nextY
		}
	}
}

func (p *Position) GetValue() string {
	return string(p.keypad[p.xAxis][p.yAxis])
}

func (p *Position) HandleMovements(movements string) {
	for _, currentMovement := range movements {
		p.Move(currentMovement)
	}
}

func part1(input []string) {
	bathroomCode := ""

	markedPosition := Position{xAxis: 1, yAxis: 1, keypad: [][]rune{
		{'1', '2', '3'},
		{'4', '5', '6'},
		{'7', '8', '9'},
	}}

	for _, steps := range input {
		markedPosition.HandleMovements(steps)
		bathroomCode = bathroomCode + markedPosition.GetValue()
	}

	fmt.Printf("Part 1: %s\n", bathroomCode)
}

func part2(input []string) {
	bathroomCode := ""

	markedPosition := Position{xAxis: 2, yAxis: 0, keypad: [][]rune{
		{'#', '#', '1', '#', '#'},
		{'#', '2', '3', '4', '#'},
		{'5', '6', '7', '8', '9'},
		{'#', 'A', 'B', 'C', '#'},
		{'#', '#', 'D', '#', '#'},
	}}

	for _, steps := range input {
		markedPosition.HandleMovements(steps)
		bathroomCode = bathroomCode + markedPosition.GetValue()
	}

	fmt.Printf("Part 2: %s\n", bathroomCode)
}

func main() {
	fileName := flag.String("file", "", "Path to the file to read")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("File name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	fileContent, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}
	input := strings.Split(string(fileContent), "\n")

	part1(input)
	part2(input)

}
