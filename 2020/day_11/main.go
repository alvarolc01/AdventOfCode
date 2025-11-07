package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	floor    rune = '.'
	empty    rune = 'L'
	occupied rune = '#'
)

var directions = [8][2]int{
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
}

func countSurrounding(currentMap []string, i, j int, marker rune) int {
	output := 0
	for _, dir := range directions {
		nx, ny := i+dir[0], j+dir[1]
		if nx < 0 || nx >= len(currentMap) || ny < 0 || ny >= len(currentMap[nx]) {
			continue
		}

		if rune(currentMap[nx][ny]) == marker {
			output++
		}
	}
	return output
}

func countVisible(currentMap []string, i, j int, marker rune) int {
	output := 0
	for _, dir := range directions {
		ci, cj := i+dir[0], j+dir[1]
		for ; ci >= 0 && ci < len(currentMap) && cj >= 0 && cj < len(currentMap[ci]); ci, cj = ci+dir[0], cj+dir[1] {
			if currentMap[ci][cj] == byte(marker) {
				output++
				break
			}

			if currentMap[ci][cj] == byte(occupied) || currentMap[ci][cj] == byte(empty) {
				break
			}
		}
	}
	return output
}

func step(grid []string, countFunc func([]string, int, int, rune) int, tolerance int) []string {
	next := make([]string, len(grid))
	for i := range grid {
		row := make([]rune, len(grid[i]))
		for j, ch := range grid[i] {
			if ch == empty && countFunc(grid, i, j, occupied) == 0 {
				row[j] = occupied
			} else if ch == occupied && countFunc(grid, i, j, occupied) >= tolerance {
				row[j] = empty
			} else {
				row[j] = ch
			}
		}
		next[i] = string(row)
	}
	return next
}

func findStatic(currentMap []string, countFunc func([]string, int, int, rune) int, tolerance int) int {
	for {
		nextMap := step(currentMap, countFunc, tolerance)
		if strings.Join(nextMap, "") == strings.Join(currentMap, "") {
			break
		}
		currentMap = nextMap
	}
	return strings.Count(strings.Join(currentMap, ""), string(occupied))
}
func part1(currentMap []string) {
	fmt.Printf("Part 1: %d\n", findStatic(currentMap, countSurrounding, 4))

}

func part2(currentMap []string) {
	fmt.Printf("Part 2: %d\n", findStatic(currentMap, countVisible, 5))
}

func main() {
	fileName := flag.String("file", "", "Path to input file")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("Use --file to specify the input file path.")
		os.Exit(1)
	}

	data, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	part1(lines)
	part2(lines)
}
