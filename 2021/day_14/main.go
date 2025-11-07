package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	stepsFirstPart  int = 10
	stepsSecondPart int = 40
)

type State struct {
	combination    string
	stepsRemaining int
}

func getFrequencies(transformations map[string]rune, combination string, stepsRemaining int, memo map[State]map[rune]int) map[rune]int {
	if stepsRemaining == 0 {
		return map[rune]int{rune(combination[0]): 1}
	}

	memKey := State{combination: combination, stepsRemaining: stepsRemaining}
	if val, ok := memo[memKey]; ok {
		return val
	}

	currTransformation := transformations[combination]
	totalFrequencies := make(map[rune]int)

	leftFrequencies := getFrequencies(transformations, string(combination[0])+string(currTransformation), stepsRemaining-1, memo)
	for key, val := range leftFrequencies {
		totalFrequencies[key] += val
	}

	rightFrequencies := getFrequencies(transformations, string(currTransformation)+string(combination[1]), stepsRemaining-1, memo)
	for key, val := range rightFrequencies {
		totalFrequencies[key] += val
	}
	memo[memKey] = totalFrequencies
	return totalFrequencies
}

func getFrequenciesAfterSteps(initialState string, transformations map[string]rune, steps int) []int {
	totalFrequencies := make(map[rune]int)
	memo := make(map[State]map[rune]int)

	for idx := 0; idx < len(initialState)-1; idx++ {
		frequencies := getFrequencies(transformations, initialState[idx:idx+2], steps, memo)
		for key, val := range frequencies {
			totalFrequencies[key] += val
		}
	}
	totalFrequencies[rune(initialState[len(initialState)-1])]++

	listFrequencies := make([]int, 0, len(totalFrequencies))
	for _, val := range totalFrequencies {
		listFrequencies = append(listFrequencies, val)
	}
	slices.Sort(listFrequencies)
	return listFrequencies
}

func part1(initialState string, transformations map[string]rune) {
	listFrequencies := getFrequenciesAfterSteps(initialState, transformations, stepsFirstPart)

	fmt.Printf("Part 1: %d\n", listFrequencies[len(listFrequencies)-1]-listFrequencies[0])
}

func part2(initialState string, transformations map[string]rune) {
	listFrequencies := getFrequenciesAfterSteps(initialState, transformations, stepsSecondPart)

	fmt.Printf("Part 2: %d\n", listFrequencies[len(listFrequencies)-1]-listFrequencies[0])
}

func parseInput(lines []string) (string, map[string]rune, error) {
	initialState := lines[0]
	transformations := make(map[string]rune)

	for _, aLine := range lines[2:] {
		parts := strings.Split(aLine, " -> ")
		transformations[parts[0]] = rune(parts[1][0])
	}

	return initialState, transformations, nil
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
	lines := strings.Split(string(fileContent), "\n")
	initialState, transformations, err := parseInput(lines)
	if err != nil {
		fmt.Println("parsing error:", err)
		os.Exit(1)
	}
	part1(initialState, transformations)
	part2(initialState, transformations)
}
