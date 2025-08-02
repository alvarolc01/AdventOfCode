package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toString(banks []int) string {
	conv := make([]string, len(banks))

	for i, n := range banks {
		conv[i] = strconv.Itoa(n)
	}

	return strings.Join(conv, "_")
}

func redistribute(banks []int) {
	maxIdx := 0
	for i := 1; i < len(banks); i++ {
		if banks[i] > banks[maxIdx] {
			maxIdx = i
		}
	}

	count := banks[maxIdx]
	banks[maxIdx] = 0

	for i := (maxIdx + 1) % len(banks); count > 0; i = (i + 1) % len(banks) {
		banks[i]++
		count--
	}
}

func part1(changes []int) {
	foundCombinations := make(map[string]bool)
	foundCombinations[toString(changes)] = true

	for found := false; !found; {
		redistribute(changes)

		if _, ok := foundCombinations[toString(changes)]; ok {
			found = true
		} else {
			foundCombinations[toString(changes)] = true
		}
	}

	fmt.Printf("Part 1: %d\n", len(foundCombinations))

}

func part2(changes []int) {
	foundCombinations := make(map[string]int)
	foundCombinations[toString(changes)] = 0
	cyclesBetweenMatch := 0

	for cycles := 1; cyclesBetweenMatch == 0; cycles++ {
		redistribute(changes)
		if val, ok := foundCombinations[toString(changes)]; ok {
			cyclesBetweenMatch = cycles - val
		} else {
			foundCombinations[toString(changes)] = cycles
		}
	}
	fmt.Printf("Part 2: %d\n", cyclesBetweenMatch)

}

func parseInput(line string) ([]int, error) {
	fields := strings.Fields(line)
	output := make([]int, len(fields))

	for idx, num := range fields {
		convertedNum, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		output[idx] = convertedNum
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

	banks, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(banks)
	part2(banks)
}
