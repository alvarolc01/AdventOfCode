package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(sortedInput [][]int) {
	checksum := 0

	for _, row := range sortedInput {
		checksum += row[len(row)-1] - row[0]
	}

	fmt.Printf("Part 1: %d\n", checksum)
}

func findEvenlyDivisedNumber(input []int) int {
	for dividendIdx := 1; dividendIdx < len(input); dividendIdx++ {
		for divisorIdx := 0; divisorIdx < dividendIdx; divisorIdx++ {
			if input[dividendIdx]%input[divisorIdx] == 0 {
				return input[dividendIdx] / input[divisorIdx]
			}
		}
	}
	return 0
}

func part2(sortedInput [][]int) {
	checksum := 0

	for _, row := range sortedInput {
		checksum += findEvenlyDivisedNumber(row)
	}

	fmt.Printf("Part 2: %d\n", checksum)
}

func parseLine(line string) []int {
	values := strings.Fields(line)
	parsedLine := make([]int, len(values))

	for idx, value := range values {
		intValue, _ := strconv.Atoi(value)
		parsedLine[idx] = intValue
	}

	return parsedLine
}

func parseInputString(input []string) [][]int {
	var parsedInput [][]int

	for _, line := range input {
		parsedInput = append(parsedInput, parseLine(line))
	}

	return parsedInput
}

func sortIndividualVectors(input [][]int) {
	for _, row := range input {
		slices.Sort(row)
	}
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

	parsedInput := parseInputString(input)
	sortIndividualVectors(parsedInput)

	part1(parsedInput)
	part2(parsedInput)
}
