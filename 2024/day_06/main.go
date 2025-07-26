package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Point struct {
	x, y int
}

type Position struct {
	Point
	direction Direction
}

type LabMap struct {
	grid [][]rune
	pos  Position
}

func (l *LabMap) Copy() *LabMap {
	newGrid := make([][]rune, len(l.grid))
	for i := range l.grid {
		newGrid[i] = make([]rune, len(l.grid[i]))
		copy(newGrid[i], l.grid[i])
	}
	return &LabMap{
		grid: newGrid,
		pos:  l.pos,
	}
}

func (l *LabMap) IsOutside(p Point) bool {
	return p.x < 0 || p.y < 0 || p.x >= len(l.grid) || p.y >= len(l.grid[0])
}

func (l *LabMap) IsWall(x, y int) bool {
	if x < 0 || y < 0 || x >= len(l.grid) || y >= len(l.grid[0]) {
		return false
	}
	return l.grid[x][y] == '#'
}

func (l *LabMap) PosInfront() (x, y int) {
	dx, dy := 0, 0
	switch l.pos.direction {
	case Up:
		dx = -1
	case Down:
		dx = 1
	case Left:
		dy = -1
	case Right:
		dy = 1
	}

	x = l.pos.x + dx
	y = l.pos.y + dy
	return
}

func (l *LabMap) GuardTakeStep() {
	nextX, nextY := l.PosInfront()

	for l.IsWall(nextX, nextY) {
		l.pos.direction = (l.pos.direction + 1) % 4
		nextX, nextY = l.PosInfront()

	}

	l.pos.x = nextX
	l.pos.y = nextY

}

func (l *LabMap) hasLoop() bool {
	visitedPositions := make(map[Position]bool)
	visitedPositions[l.pos] = true

	for !l.IsOutside(l.pos.Point) {
		l.GuardTakeStep()
		if _, ok := visitedPositions[l.pos]; ok {
			return true
		}
		visitedPositions[l.pos] = true
	}
	return false

}

func NewLabMap(coordinates []string) *LabMap {
	grid := make([][]rune, len(coordinates))
	var startPosition *Position

	for r, row := range coordinates {
		grid[r] = []rune(row)
		if startPosition == nil {
			for c, char := range grid[r] {
				if char == '^' {
					startPosition = &Position{
						Point:     Point{x: r, y: c},
						direction: Up,
					}
					grid[r][c] = '.'
					break
				}
			}
		}
	}

	return &LabMap{
		grid: grid,
		pos:  *startPosition,
	}
}

func part1(lab *LabMap) map[Point]bool {
	visitedPositions := make(map[Point]bool)

	for !lab.IsOutside(lab.pos.Point) {
		visitedPositions[lab.pos.Point] = true
		lab.GuardTakeStep()
	}

	fmt.Printf("Part 1: %d\n", len(visitedPositions))
	return visitedPositions
}

func part2(lab *LabMap, visitedPositions map[Point]bool) {
	possibleLoops := 0
	startingPos := lab.pos

	for point := range visitedPositions {
		if point.x == startingPos.x && point.y == startingPos.y {
			continue
		} else if lab.IsOutside(point) {
			continue
		}

		if lab.grid[point.x][point.y] == '.' {
			lab.grid[point.x][point.y] = '#'

			lab.pos = startingPos
			if lab.hasLoop() {
				possibleLoops++
			}

			lab.grid[point.x][point.y] = '.'
		}
	}

	fmt.Printf("Part 2: %d\n", possibleLoops)
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

	lab := NewLabMap(input)
	if lab == nil {
		fmt.Println("error parsing input: ", err)
		os.Exit(1)
	}

	copyLab := lab.Copy()
	visitedPos := part1(lab)

	part2(copyLab, visitedPos)
}
