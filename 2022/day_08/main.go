package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction struct {
	dx, dy int
}

var directions = []Direction{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func withinBounds(forest [][]int, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(forest) && y < len(forest[0])
}

func isVisible(forest [][]int, i, j int) bool {
	currHeight := forest[i][j]

	for _, dir := range directions {
		visible := true

		for x, y := i+dir.dx, j+dir.dy; withinBounds(forest, x, y); x, y = x+dir.dx, y+dir.dy {
			if forest[x][y] >= currHeight {
				visible = false
				break
			}
		}

		if visible {
			return true
		}
	}
	return false
}

func part1(forestMap [][]int) {
	visibleTrees := 0

	for i := range forestMap {
		for j := range forestMap[i] {
			if isVisible(forestMap, i, j) {
				visibleTrees++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", visibleTrees)
}

func calculateScenicScore(forest [][]int, i, j int) int {
	currHeight := forest[i][j]
	score := 1

	for _, d := range directions {
		count := 0
		for x, y := i+d.dx, j+d.dy; withinBounds(forest, x, y); x, y = x+d.dx, y+d.dy {
			count++
			if forest[x][y] >= currHeight {
				break
			}
		}
		score *= count
	}

	return score
}

func part2(forestMap [][]int) {
	maxScenicScore := 0

	for i := range forestMap {
		for j := range forestMap[i] {
			if currScore := calculateScenicScore(forestMap, i, j); currScore > maxScenicScore {
				maxScenicScore = currScore
			}
		}
	}

	fmt.Printf("Part 2: %d\n", maxScenicScore)
}

func parseInput(lines []string) ([][]int, error) {
	result := make([][]int, len(lines))
	for idx := range result {
		result[idx] = make([]int, len(lines[0]))
	}

	for row, line := range lines {
		for col, char := range line {
			cnvNum, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}

			result[row][col] = cnvNum
		}
	}

	return result, nil
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
	forestMap, err := parseInput(input)
	if err != nil {
		fmt.Println("error reading input file:", err)
		os.Exit(1)
	}

	part1(forestMap)
	part2(forestMap)
}
