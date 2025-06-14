package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func part1(leftCol, rightCol []int) {
	slices.Sort(leftCol)
	slices.Sort(rightCol)

	totalDistance := 0
	for idx, val := range leftCol {
		totalDistance += int(math.Abs(float64(val - rightCol[idx])))
	}

	fmt.Printf("Part 1: %d\n", totalDistance)
}

func part2(leftCol, rightCol []int) {
	repetitionsRight := make(map[int]int)
	for _, val := range rightCol {
		repetitionsRight[val]++
	}

	similarityScore := 0
	for _, key := range leftCol {
		similarityScore += repetitionsRight[key] * key
	}

	fmt.Printf("Part 2: %d\n", similarityScore)
}

func parseInput(input string) (leftCol, rightCol []int, err error) {
	lines := strings.Split(input, "\n")

	leftCol = make([]int, len(lines))
	rightCol = make([]int, len(lines))

	for idx, line := range lines {
		var leftNum, rightNum int
		_, err = fmt.Sscanf(line, "%d\t%d", &leftNum, &rightNum)
		if err != nil {
			return
		}

		leftCol[idx] = leftNum
		rightCol[idx] = rightNum
	}

	return
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

	leftCol, rightCol, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(leftCol, rightCol)
	part2(leftCol, rightCol)
}
