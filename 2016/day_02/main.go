package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func readFile(fileName *string) []string {
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read file content: %s", err)
		os.Exit(1)
	}

	return strings.Split(string(content), "\n")

}

type Position struct {
	xAxis  int
	yAxis  int
	keypad [][]rune
}

func (p *Position) Move(direction rune) {
	switch direction {
	case 'L':
		p.MoveLeft()
	case 'R':
		p.MoveRight()
	case 'U':
		p.MoveUp()
	case 'D':
		p.MoveDown()
	}
}

func (p *Position) MoveLeft() {
	if p.yAxis > 0 && p.keypad[p.xAxis][p.yAxis-1] != '#' {
		p.yAxis--
	}
}

func (p *Position) MoveRight() {
	if p.yAxis < len(p.keypad)-1 && p.keypad[p.xAxis][p.yAxis+1] != '#' {
		p.yAxis++
	}
}

func (p *Position) MoveUp() {
	if p.xAxis > 0 && p.keypad[p.xAxis-1][p.yAxis] != '#' {
		p.xAxis--
	}
}

func (p *Position) MoveDown() {
	if p.xAxis < len(p.keypad)-1 && p.keypad[p.xAxis+1][p.yAxis] != '#' {
		p.xAxis++
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

	inputString := readFile(fileName)

	part1(inputString)
	part2(inputString)

}
