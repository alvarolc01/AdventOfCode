package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

func isMinimum(board []string, i, j int) bool {
	h := board[i][j]
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range dirs {
		ni, nj := i+dir[0], j+dir[1]
		if ni >= 0 && nj >= 0 && ni < len(board) && nj < len(board[0]) && board[ni][nj] <= h {
			return false
		}
	}
	return true
}

func part1(board []string) {
	sumRiskLevels := 0

	for i, row := range board {
		for j, col := range row {
			if isMinimum(board, i, j) {
				sumRiskLevels += int(col-'0') + 1
			}
		}
	}

	fmt.Printf("Part 1: %d\n", sumRiskLevels)
}

type Point struct {
	x, y int
}

func getBasinSize(board []string, i, j int) int {
	pointsBasin := make(map[Point]bool)
	q := []Point{{i, j}}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if _, ok := pointsBasin[curr]; ok {
			continue
		}
		if curr.x < 0 || curr.y < 0 || curr.x >= len(board) || curr.y >= len(board[0]) || board[curr.x][curr.y] == '9' {
			continue
		}
		pointsBasin[curr] = true
		q = append(q, Point{curr.x, curr.y + 1}, Point{curr.x, curr.y - 1}, Point{curr.x + 1, curr.y}, Point{curr.x - 1, curr.y})
	}

	return len(pointsBasin)
}

func part2(board []string) {
	basinSizes := []int{}

	for i, row := range board {
		for j, _ := range row {
			if isMinimum(board, i, j) {
				basinSize := getBasinSize(board, i, j)
				basinSizes = append(basinSizes, basinSize)
			}
		}
	}

	sort.Ints(basinSizes)

	ans := 1
	for i := 0; i < 3; i++ {
		ans *= basinSizes[len(basinSizes)-1-i]
	}
	fmt.Printf("Part 2: %d\n", ans)
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

	part1(input)
	part2(input)
}
