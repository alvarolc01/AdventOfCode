package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

type Line struct {
	start, end Point
}

func (l *Line) isHorizontalOrVertical() bool {
	return l.start.x == l.end.x || l.start.y == l.end.y
}

func (l *Line) isDiagonal() bool {
	return math.Abs(float64(l.start.x-l.end.x)) == math.Abs(float64(l.start.y-l.end.y))
}

func NewLine(line string) (*Line, error) {
	var startX, startY, endX, endY int
	_, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &startX, &startY, &endX, &endY)
	if err != nil {
		return nil, err
	}

	return &Line{
		start: Point{x: startX, y: startY},
		end:   Point{x: endX, y: endY},
	}, nil
}

func step(start, end int) int {
	if start == end {
		return 0
	} else if start > end {
		return -1
	}
	return 1
}

func addLinePoints(points map[Point]int, line *Line) {
	dx := step(line.start.x, line.end.x)
	dy := step(line.start.y, line.end.y)

	for curr := line.start; ; curr.x, curr.y = curr.x+dx, curr.y+dy {
		points[curr]++
		if curr == line.end {
			break
		}
	}
}

func countOverlappedPoints(points map[Point]int) int {
	count := 0
	for _, val := range points {
		if val >= 2 {
			count++
		}
	}
	return count
}

func part1(lines []*Line) {
	countPoints := make(map[Point]int)

	for _, line := range lines {
		if line.isHorizontalOrVertical() {
			addLinePoints(countPoints, line)
		}
	}

	overlappingPoints := countOverlappedPoints(countPoints)
	fmt.Printf("Part 1: %d\n", overlappingPoints)
}

func part2(lines []*Line) {
	countPoints := make(map[Point]int)

	for _, line := range lines {
		if line.isHorizontalOrVertical() || line.isDiagonal() {
			addLinePoints(countPoints, line)
		}
	}

	overlappingPoints := countOverlappedPoints(countPoints)
	fmt.Printf("Part 2: %d\n", overlappingPoints)
}

func parseInput(lines []string) ([]*Line, error) {
	output := make([]*Line, len(lines))
	for idx, line := range lines {
		parsedLine, err := NewLine(line)
		if err != nil {
			return nil, err
		}
		output[idx] = parsedLine
	}

	return output, nil
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

	lines, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(lines)
	part2(lines)
}
