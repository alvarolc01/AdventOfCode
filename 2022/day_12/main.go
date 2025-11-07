package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	currHeight int
	distance   int
	x, y       int
	index      int
}

var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func getInitialPositions(initialMap []string, marker rune) (int, int) {
	for i, row := range initialMap {
		for j, cell := range row {
			if cell == marker {
				return i, j
			}
		}
	}
	return -1, -1
}

func getHeight(ch rune) int {
	height := int(ch - 'a')
	if ch == 'S' {
		height = 0
	} else if ch == 'E' {
		height = 25
	}

	return height
}

func bfs(initialMap []string, iniX, iniY int, target rune, validStep func(int, int) bool) int {
	rows, cols := len(initialMap), len(initialMap[0])
	queue := []*Node{{currHeight: getHeight(rune(initialMap[iniX][iniY])), x: iniX, y: iniY, distance: 0}}

	found := make(map[[2]int]bool)

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if found[[2]int{top.x, top.y}] {
			continue
		}
		found[[2]int{top.x, top.y}] = true

		if rune(initialMap[top.x][top.y]) == target {
			return top.distance
		}

		for _, dir := range directions {
			nx, ny := top.x+dir[0], top.y+dir[1]
			if nx < 0 || ny < 0 || nx >= rows || ny >= cols {
				continue
			}

			nextHeight := getHeight(rune(initialMap[nx][ny]))
			if !validStep(nextHeight, top.currHeight) {
				continue
			}

			queue = append(queue, &Node{
				distance:   top.distance + 1,
				x:          nx,
				y:          ny,
				currHeight: nextHeight,
			})
		}
	}

	return -1
}

func part1(initialMap []string) {
	initialX, initialY := getInitialPositions(initialMap, 'S')
	heightCheck := func(next, curr int) bool {
		return next-1 <= curr
	}
	fmt.Printf("Part 1: %d\n", bfs(initialMap, initialX, initialY, 'E', heightCheck))
}

func part2(initialMap []string) {
	initialX, initialY := getInitialPositions(initialMap, 'E')
	heightCheck := func(next, curr int) bool {
		return next >= curr-1
	}
	fmt.Printf("Part 2: %d\n", bfs(initialMap, initialX, initialY, 'a', heightCheck))
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
	lines := strings.Split(string(fileContent), "\n")

	part1(lines)
	part2(lines)
}
