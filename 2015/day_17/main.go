package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const EggnogLiters int = 150

func backtrackingCombinations(options, validCombinationsByLength []int, currentSpace, position, currentContainers int) {
	if currentSpace == EggnogLiters {
		validCombinationsByLength[currentContainers]++
		return
	}
	if currentSpace > EggnogLiters || position >= len(options) {
		return
	}

	backtrackingCombinations(options, validCombinationsByLength, currentSpace+options[position], position+1, currentContainers+1)
	backtrackingCombinations(options, validCombinationsByLength, currentSpace, position+1, currentContainers)
}

func part1(input []int) {
	validCombinationsByLength := make([]int, len(input)+1)
	backtrackingCombinations(input, validCombinationsByLength, 0, 0, 0)

	totalCombinations := 0
	for _, combinations := range validCombinationsByLength {
		totalCombinations += combinations
	}

	fmt.Printf("Part 1: %d\n", totalCombinations)
}

func part2(input []int) {
	validCombinationsByLength := make([]int, len(input)+1)
	backtrackingCombinations(input, validCombinationsByLength, 0, 0, 0)

	combinationsMinNumContainers := 0
	for _, combinations := range validCombinationsByLength {
		if combinationsMinNumContainers == 0 {
			combinationsMinNumContainers = combinations
		}
	}

	fmt.Printf("Part 2: %d\n", combinationsMinNumContainers)
}

func parseInput(input []string) []int {
	var result []int
	for _, line := range input {
		convertedNum, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		result = append(result, convertedNum)
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
