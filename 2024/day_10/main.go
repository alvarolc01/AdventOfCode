package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

var dirs = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func getReachableTops(topMap []string, i, j int) map[Point]int {
	visitedTops := make(map[Point]int)

	q := []Point{{i, j}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if topMap[curr.x][curr.y] == '9' {
			visitedTops[curr]++
			continue
		}

		for _, dir := range dirs {
			next := curr
			next.x += dir[0]
			next.y += dir[1]

			if next.x < 0 || next.x >= len(topMap) || next.y < 0 || next.y >= len(topMap[0]) {
				continue
			} else if topMap[next.x][next.y] != topMap[curr.x][curr.y]+1 {
				continue
			}
			q = append(q, next)
		}
	}

	return visitedTops
}

func part1(topMap []string) {
	output := 0

	for i, row := range topMap {
		for j, ch := range row {
			if ch == '0' {
				reachableTops := getReachableTops(topMap, i, j)
				output += len(reachableTops)
			}
		}
	}

	fmt.Printf("Part 1: %d\n", output)
}

func part2(topMap []string) {
	output := 0

	for i, row := range topMap {
		for j, ch := range row {
			if ch == '0' {
				reachableTops := getReachableTops(topMap, i, j)
				for _, routes := range reachableTops {
					output += routes
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", output)
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
	topMap := strings.Split(string(fileContent), "\n")

	part1(topMap)
	part2(topMap)
}
