package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	TimesFirstPart  int = 40
	TimesSecondPart int = 50
)

func convertArray(input []int) []int {
	var result []int

	numRepetitions := 1
	currentValue := input[0]
	for idx := 1; idx < len(input); idx++ {
		if input[idx] == currentValue {
			numRepetitions++
		} else {
			result = append(result, numRepetitions)
			result = append(result, currentValue)
			numRepetitions = 1
			currentValue = input[idx]
		}
	}

	result = append(result, numRepetitions)
	result = append(result, currentValue)

	return result
}

func part1(input []int) {
	for i := 0; i < TimesFirstPart; i++ {
		input = convertArray(input)
	}

	fmt.Printf("Part 1: %d\n", len(input))
}

func part2(input []int) {
	for i := 0; i < TimesSecondPart; i++ {
		input = convertArray(input)
	}

	fmt.Printf("Part 2: %d\n", len(input))
}

func convertInput(input string) []int {
	var result []int
	for _, char := range input {
		result = append(result, int(char-'0'))
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
	input := string(fileContent)

	convertedInput := convertInput(input)

	part1(convertedInput)
	part2(convertedInput)
}
