package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	sideLength        int = 71
	initialCorruption int = 1024
)

type Position struct {
	x, y int
}

type Node struct {
	pos  Position
	dist int
}

var dirs = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func findDistance(corruptedPositions map[Position]bool) int {
	visited := map[Position]bool{}
	q := []Node{{pos: Position{0, 0}, dist: 0}}

	for len(q) > 0 {
		top := q[0]
		q = q[1:]

		if top.pos.x < 0 || top.pos.y < 0 || top.pos.x >= sideLength || top.pos.y >= sideLength {
			continue
		}
		if visited[top.pos] {
			continue
		}
		if corruptedPositions[top.pos] {
			continue
		}
		if top.pos.x == sideLength-1 && top.pos.y == sideLength-1 {
			return top.dist

		}
		visited[top.pos] = true

		for _, d := range dirs {
			nx, ny := top.pos.x+d[0], top.pos.y+d[1]
			q = append(q,
				Node{
					pos:  Position{nx, ny},
					dist: top.dist + 1,
				})
		}
	}
	return -1
}

func part1(corruptedPositions []Position) {
	mapCorruptedPositions := make(map[Position]bool)
	for i := 0; i < initialCorruption; i++ {
		mapCorruptedPositions[corruptedPositions[i]] = true
	}
	output := findDistance(mapCorruptedPositions)

	fmt.Printf("Part 1: %d\n", output)
}

func part2(corruptedPositions []Position) {
	left, right := initialCorruption, len(corruptedPositions)-1
	var output Position

	for left <= right {
		mid := (left + right) / 2
		mapCorruptedPositions := make(map[Position]bool)
		for i := 0; i <= mid; i++ {
			mapCorruptedPositions[corruptedPositions[i]] = true
		}

		distance := findDistance(mapCorruptedPositions)
		if distance == -1 {
			output = corruptedPositions[mid]
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	fmt.Printf("Part 2: %s\n", fmt.Sprintf("%d,%d", output.x, output.y))
}

func parseInput(lines []string) ([]Position, error) {
	output := make([]Position, 0, len(lines))
	for _, line := range lines {
		var x, y int
		_, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil {
			return nil, err
		}
		output = append(output, Position{x, y})
	}
	return output, nil
}

func main() {
	fileName := flag.String("file", "", "Path to the file")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("file name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	content, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("error reading input file:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")
	corruptedPositions, err := parseInput(lines)
	if err != nil {
		fmt.Println("parsing error:", err)
		os.Exit(1)
	}

	part1(corruptedPositions)
	part2(corruptedPositions)
}
