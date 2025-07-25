package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(changes []int) {
	finalFrequency := 0

	for _, change := range changes {
		finalFrequency += change
	}

	fmt.Printf("Part 1: %d\n", finalFrequency)

}

func part2(changes []int) {
	foundFrequencies := make(map[int]bool)
	currFreq := 0
	found := false

	for idx := 0; !found; idx = (idx + 1) % len(changes) {
		currFreq += changes[idx]
		found = foundFrequencies[currFreq]
		foundFrequencies[currFreq] = true
	}

	fmt.Printf("Part 2: %d\n", currFreq)

}

func parseInput(lines []string) ([]int, error) {
	output := make([]int, len(lines))

	for idx, line := range lines {
		isNegative := line[0] == '-'
		claim, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}

		if isNegative {
			claim *= -1
		}

		output[idx] = claim
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

	changes, err := parseInput(input)
	if err != nil {
		fmt.Println("error creating graph:", err)
		os.Exit(1)
	}

	part1(changes)
	part2(changes)
}
