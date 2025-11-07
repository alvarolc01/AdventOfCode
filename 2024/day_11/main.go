package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	blinksFirstPart  int = 25
	blinksSecondPart int = 75
)

type SeenStone struct {
	depth int
	num   int
}

func calculateStonesAfterBlinks(seen map[SeenStone]int, stone, remainingBlinks int) int {
	if remainingBlinks == 0 {
		return 1
	}
	seenStoneKey := SeenStone{remainingBlinks, stone}
	if val, ok := seen[seenStoneKey]; ok {
		return val
	}

	var nextVals []int
	if stone == 0 {
		nextVals = append(nextVals, 1)

	} else if len(strconv.Itoa(stone))%2 == 0 {
		currStone := strconv.Itoa(stone)
		mid := len(currStone) / 2

		firstStone, _ := strconv.Atoi(currStone[:mid])
		secondStone, _ := strconv.Atoi(currStone[mid:])

		nextVals = []int{firstStone, secondStone}
	} else {
		nextVals = append(nextVals, stone*2024)
	}

	numStones := 0
	for _, num := range nextVals {
		numStones += calculateStonesAfterBlinks(seen, num, remainingBlinks-1)
	}

	seen[seenStoneKey] = numStones
	return numStones
}

func part1(stones []int) {
	totalStones := 0
	seen := make(map[SeenStone]int)
	for _, aStone := range stones {
		totalStones += calculateStonesAfterBlinks(seen, aStone, blinksFirstPart)
	}

	fmt.Printf("Part 1: %d\n", totalStones)
}

func part2(stones []int) {
	totalStones := 0
	seen := make(map[SeenStone]int)
	for _, aStone := range stones {
		totalStones += calculateStonesAfterBlinks(seen, aStone, blinksSecondPart)
	}

	fmt.Printf("Part 2: %d\n", totalStones)
}

func parseInput(line string) ([]int, error) {
	fields := strings.Fields(line)
	output := make([]int, 0, len(fields))

	for _, aStone := range fields {
		stoneVal, err := strconv.Atoi(aStone)
		if err != nil {
			return nil, err
		}
		output = append(output, stoneVal)
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
	input := string(fileContent)
	stones, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(stones)
	part2(stones)
}
