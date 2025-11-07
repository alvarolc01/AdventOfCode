package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bfs(items map[int][]int, start int, found map[int]bool) {
	q := []int{start}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if _, ok := found[curr]; ok {
			continue
		}
		found[curr] = true

		for _, connectedNode := range items[curr] {
			q = append(q, connectedNode)
		}
	}
}

func part1(items map[int][]int) {
	found := make(map[int]bool)
	bfs(items, 0, found)
	fmt.Printf("Part 1: %d\n", len(found))

}

func part2(items map[int][]int) {
	found := make(map[int]bool)
	islands := 0
	for id := range items {
		if !found[id] {
			islands++
			bfs(items, id, found)
		}
	}

	fmt.Printf("Part 2: %d\n", islands)

}

func parseInput(lines []string) (map[int][]int, error) {
	output := make(map[int][]int)

	for _, aLine := range lines {
		aLine = strings.TrimSpace(aLine)

		parts := strings.Split(aLine, "<->")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid format")
		}

		id, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid id: %v", err)
		}

		followingNums := strings.Split(parts[1], ",")
		neighbours := make([]int, 0, len(followingNums))

		for _, num := range followingNums {
			convNum, err := strconv.Atoi(strings.TrimSpace(num))
			if err != nil {
				return nil, err
			}

			neighbours = append(neighbours, convNum)
		}

		output[id] = neighbours

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

	program, err := parseInput(input)
	if err != nil {
		fmt.Println("error parting input:", err)
		os.Exit(1)
	}

	part1(program)
	part2(program)
}
