package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func updateInput(input []int, currentPosition int, offsetUpdate func(int) int) int {
	previousPosition := currentPosition
	currentPosition += input[currentPosition]
	input[previousPosition] += offsetUpdate(input[previousPosition])

	return currentPosition
}

func calculateSteps(input []int, offsetUpdate func(int) int) int {
	numSteps := 0
	currentPosition := 0

	for currentPosition >= 0 && currentPosition < len(input) {
		numSteps++
		currentPosition = updateInput(input, currentPosition, offsetUpdate)
	}

	return numSteps
}

func part1(input []int) {
	offsetUpdate := func(currentOffset int) int {
		return 1
	}
	numSteps := calculateSteps(input, offsetUpdate)

	fmt.Printf("Part 1: %d\n", numSteps)
}

func part2(input []int) {
	offsetUpdate := func(currentOffset int) int {
		if currentOffset >= 3 {
			return -1
		}
		return 1
	}
	numSteps := calculateSteps(input, offsetUpdate)

	fmt.Printf("Part 2: %d\n", numSteps)
}

func parseInputToInt(input []string) []int {
	var parsedInput []int

	for _, value := range input {
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("Failed to convert string to int: %s\n", err)
			os.Exit(1)
		}
		parsedInput = append(parsedInput, valueInt)
	}

	return parsedInput
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

	parsedInput := parseInputToInt(input)
	parsedInputCopy := make([]int, len(parsedInput))
	copy(parsedInputCopy, parsedInput)

	part1(parsedInput)
	part2(parsedInputCopy)

}
