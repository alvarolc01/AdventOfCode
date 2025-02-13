package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	North = iota
	East
	South
	West
)

type Direction int

func (d *Direction) TurnLeft()  { *d = (*d - 1 + 4) % 4 }
func (d *Direction) TurnRight() { *d = (*d + 1) % 4 }

var directionVectors = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
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

	directionX := directionVectors[p.currentDirection][0]
	directionY := directionVectors[p.currentDirection][1]

	p.currentPoint.x += directionX * stepsToMove
	p.currentPoint.y += directionY * stepsToMove

}

func (p *Position) ProcessCommandFirstVisited(command string) bool {
	directionRune, stepsToMove, err := p.ParseCommand(command)
	if err != nil {
		return false
	}

	p.ChangeDirection(directionRune)

	directionX := directionVectors[p.currentDirection][0]
	directionY := directionVectors[p.currentDirection][1]

	for i := 0; i < stepsToMove; i++ {
		p.currentPoint.x += directionX
		p.currentPoint.y += directionY

		if _, ok := p.visited[p.currentPoint]; ok {
			return true
		}
		p.visited[p.currentPoint] = true
	}

	return false
}

func NewPosition() *Position {
	return &Position{
		currentPoint:     Point{x: 0, y: 0},
		currentDirection: North,
		visited:          make(map[Point]bool),
	}
}

func part1(input string) {
	commands := strings.Split(input, ", ")
	currentPos := NewPosition()

	for _, command := range commands {
		currentPos.ProcessCommand(command)
	}

	fmt.Printf("Part 1: %d\n", currentPos.currentPoint.Distance())
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

	fmt.Printf("Part 2: %d\n", currentPos.currentPoint.Distance())
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

	part1(input)
	part2(input)

}
