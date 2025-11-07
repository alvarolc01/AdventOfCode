package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Position struct {
	x, y int
}

type Node struct {
	pos      Position
	distance int
	found    map[int]bool
}

type SearchState struct {
	pos   Position
	found string
}

var directions = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func getStartingData(initialMap []string) (Position, int) {
	var startingPos Position
	countNumbers := 0
	for x, row := range initialMap {
		for y, cell := range row {
			if cell == '0' {
				startingPos = Position{x, y}
			}
			if cell != '#' && cell != '.' {
				countNumbers++
			}
		}
	}
	return startingPos, countNumbers
}

func getStateKey(curr *Node) SearchState {
	currFound := make([]rune, 0, len(curr.found))
	for key := range curr.found {
		currFound = append(currFound, rune(key+'0'))
	}
	slices.Sort(currFound)
	return SearchState{
		pos:   curr.pos,
		found: string(currFound),
	}
}

func bfs(initialMap []string, mustReturnToOrigin bool) int {
	initialPos, totalNumbers := getStartingData(initialMap)
	q := []Node{{
		pos:      initialPos,
		distance: 0,
		found:    map[int]bool{0: true},
	}}
	mem := make(map[SearchState]bool)

	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		cell := initialMap[current.pos.x][current.pos.y]
		if cell == '#' {
			continue
		} else if cell != '.' {
			current.found[int(cell-'0')] = true
		}

		if len(current.found) == totalNumbers && (!mustReturnToOrigin || cell == '0') {
			return current.distance
		}

		stateKey := getStateKey(&current)
		if mem[stateKey] {
			continue
		}
		mem[stateKey] = true

		for _, dir := range directions {
			nx, ny := current.pos.x+dir[0], current.pos.y+dir[1]
			if nx < 0 || ny < 0 || nx >= len(initialMap) || ny >= len(initialMap[nx]) {
				continue
			}
			nextFound := make(map[int]bool, len(current.found))
			for k, v := range current.found {
				nextFound[k] = v
			}

			q = append(q, Node{
				pos:      Position{nx, ny},
				distance: current.distance + 1,
				found:    nextFound,
			})

		}
	}

	return -1
}

func part1(initialMap []string) {
	fmt.Printf("Part 1: %d\n", bfs(initialMap, false))
}

func part2(initialMap []string) {
	fmt.Printf("Part 2: %d\n", bfs(initialMap, true))
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

	initialMap := strings.Split(string(fileContent), "\n")

	part1(initialMap)
	part2(initialMap)
}
