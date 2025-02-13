package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	NumGroupsFirstPart  int = 3
	NumGroupsSecondPart int = 4
)

func backtrackingFinder(values []int, pos, target, packagesQE, accumulatedWeight, numSelected int) (int, int) {
	if accumulatedWeight == target {
		return packagesQE, numSelected
	} else if pos >= len(values) || accumulatedWeight > target {
		return math.MaxInt, math.MaxInt
	}

	takingCurrentQE, takingCurrentRequiredCount := backtrackingFinder(values, pos+1, target,
		packagesQE*values[pos], accumulatedWeight+values[pos], numSelected+1)

	notTakingCurrentQE, notTakingCurrentRequiredCount := backtrackingFinder(values, pos+1, target,
		packagesQE, accumulatedWeight, numSelected)

	if takingCurrentRequiredCount > notTakingCurrentRequiredCount {
		return notTakingCurrentQE, notTakingCurrentRequiredCount
	} else if takingCurrentRequiredCount < notTakingCurrentRequiredCount {
		return takingCurrentQE, takingCurrentRequiredCount
	}

	if takingCurrentQE < notTakingCurrentQE {
		return takingCurrentQE, takingCurrentRequiredCount
	}
	return notTakingCurrentQE, notTakingCurrentRequiredCount
}

func part1(packages []int) {
	total := 0
	for _, value := range packages {
		total += value
	}
	target := total / NumGroupsFirstPart
	qe, _ := backtrackingFinder(packages, 0, target, 1, 0, 0)
	fmt.Printf("Part 1: %d\n", qe)
}

func part2(packages []int) {
	total := 0
	for _, value := range packages {
		total += value
	}
	target := total / NumGroupsSecondPart
	qe, _ := backtrackingFinder(packages, 0, target, 1, 0, 0)
	fmt.Printf("Part 2: %d\n", qe)
}

func parseInput(input []string) []int {
	var result []int
	for _, line := range input {
		num, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		result = append(result, num)
	}
	return result
}

func main() {
	fileName := flag.String("file", "", "Path to the file to read")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("File name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	fileContent, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}
	input := strings.Split(string(fileContent), "\n")

	parsedInput := parseInput(input)

	part1(parsedInput)
	part2(parsedInput)
}
