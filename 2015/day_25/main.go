package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	stepMultiplier int = 252533
	stepDivider    int = 33554393
	initialValue   int = 20151125
)

type Point struct {
	row, col int
}

func (p *Point) moveToNextPoint() {
	if p.row == 1 {
		p.row = p.row + p.col
		p.col = 1
	} else {
		p.row--
		p.col++
	}
}

func step(currentValue int) int {
	return (currentValue * stepMultiplier) % stepDivider
}

func part1(row, column int) {
	currentPoint := Point{1, 1}
	currentValue := initialValue

	targetPoint := Point{row, column}

	for currentPoint != targetPoint {
		currentPoint.moveToNextPoint()
		currentValue = step(currentValue)
	}

	fmt.Printf("Part 1 %d\n", currentValue)
}

func getPosition(input string) (row, col int) {
	re := regexp.MustCompile("(\\d+)")
	matches := re.FindAllString(input, -1)

	row, errRow := strconv.Atoi(matches[0])
	col, errCol := strconv.Atoi(matches[1])

	if errRow != nil || errCol != nil {
		fmt.Println("Invalid position provided")
	}

	return
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
	input := string(fileContent)

	row, col := getPosition(input)

	part1(row, col)
}
