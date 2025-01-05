package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(fileName *string) []string {
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read file content: %s", err)
		os.Exit(1)
	}

	return strings.Split(string(content), "\n")

}

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
	var parsedLine []int
	values := strings.Split(line, "\t")
	for _, value := range values {
		intValue, _ := strconv.Atoi(value)
		parsedLine = append(parsedLine, intValue)
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

	inputString := readFile(fileName)
	parsedInput := parseInputString(inputString)
	sortIndividualVectors(parsedInput)

	part1(parsedInput)
	part2(parsedInput)

}
