package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFile(fileName *string) string {
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

	return string(content)

}

type Direction struct {
	directionNum int
}

func (d *Direction) TurnLeft() {
	d.directionNum = (d.directionNum - 1 + 4) % 4
}

func (d *Direction) TurnRight() {
	d.directionNum = (d.directionNum + 1) % 4
}

func (d *Direction) AdvanceVector() (int, int) {
	switch d.directionNum {
	case 0:
		return 0, 1
	case 1:
		return 1, 0
	case 2:
		return 0, -1
	case 3:
		return -1, 0
	}

	return 0, 0
}

type Point struct {
	x int
	y int
}

func (p *Point) Distance() int {
	return int(math.Abs(float64(p.x))) + int(math.Abs(float64(p.y)))
}

type Position struct {
	currentPoint     Point
	currentDirection Direction
	visited          map[Point]bool
}

func (p *Position) ChangeDirection(direction rune) {
	if direction == 'L' {
		p.currentDirection.TurnLeft()
	} else if direction == 'R' {
		p.currentDirection.TurnRight()
	}
}

func (p *Position) ParseCommand(command string) (rune, int, error) {
	directionRune := rune(command[0])

	stepsToMove, err := strconv.Atoi(command[1:])
	if err != nil {
		return ' ', 0, err
	}

	return directionRune, stepsToMove, nil
}

func (p *Position) ProcessCommand(command string) {
	directionRune, stepsToMove, err := p.ParseCommand(command)
	if err != nil {
		return
	}

	p.ChangeDirection(directionRune)

	moveMultX, moveMultY := p.currentDirection.AdvanceVector()

	p.currentPoint.x += moveMultX * stepsToMove
	p.currentPoint.y += moveMultY * stepsToMove

}

func (p *Position) ProcessCommandFirstVisited(command string) bool {
	directionRune, stepsToMove, err := p.ParseCommand(command)
	if err != nil {
		return false
	}

	p.ChangeDirection(directionRune)

	moveMultX, moveMultY := p.currentDirection.AdvanceVector()

	for i := 0; i < stepsToMove; i++ {
		p.currentPoint.x += moveMultX
		p.currentPoint.y += moveMultY

		if _, ok := p.visited[p.currentPoint]; ok {
			return true
		}
		p.visited[p.currentPoint] = true
	}

	return false
}

func (p *Position) CalculateDistance() int {
	return p.currentPoint.Distance()
}

func NewPosition() *Position {
	return &Position{
		currentPoint:     Point{x: 0, y: 0},
		currentDirection: Direction{directionNum: 0},
		visited:          make(map[Point]bool),
	}
}

func part1(input string) {
	commands := strings.Split(input, ", ")
	currentPos := NewPosition()

	for _, command := range commands {
		currentPos.ProcessCommand(command)
	}

	fmt.Printf("Part 1: %d\n", currentPos.CalculateDistance())
}

func part2(input string) {
	commands := strings.Split(input, ", ")
	currentPos := NewPosition()

	for _, command := range commands {
		visitedPoint := currentPos.ProcessCommandFirstVisited(command)
		if visitedPoint {
			break
		}
	}

	fmt.Printf("Part 2: %d\n", currentPos.CalculateDistance())
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
