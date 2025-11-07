package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

var (
	minTowelLength int = math.MaxInt
	maxTowelLength int = math.MinInt
)

func possibleCombinations(towels map[string]bool, targetDesign string, idx int, mem map[int]int) int {
	if len(targetDesign) == idx {
		return 1
	}

	if val, ok := mem[idx]; ok {
		return val
	}

	totalCombinations := 0
	for i := minTowelLength; i <= maxTowelLength && len(targetDesign) >= idx+i; i++ {
		requiredDesign := targetDesign[idx : idx+i]
		if _, ok := towels[requiredDesign]; ok {
			totalCombinations += possibleCombinations(towels, targetDesign, idx+i, mem)
		}
	}
	mem[idx] = totalCombinations

	return totalCombinations
}

func part1(towels map[string]bool, designs []string) {
	result := 0

	for _, aDesign := range designs {
		mem := make(map[int]int)
		if possibleCombinations(towels, aDesign, 0, mem) != 0 {
			result++
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func part2(towels map[string]bool, designs []string) {
	result := 0

	for _, aDesign := range designs {
		mem := make(map[int]int)
		result += possibleCombinations(towels, aDesign, 0, mem)
	}

	fmt.Printf("Part 2: %d\n", result)
}

func parseInput(lines []string) (map[string]bool, []string, error) {
	availableTowels := make(map[string]bool)
	towels := strings.Split(lines[0], ",")
	for _, aTowel := range towels {
		towelName := strings.TrimSpace(aTowel)
		minTowelLength = min(minTowelLength, len(towelName))
		maxTowelLength = max(maxTowelLength, len(towelName))
		availableTowels[towelName] = true
	}

	return availableTowels, lines[2:], nil
}
func main() {
	fileName := flag.String("file", "", "Path to the file")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("file name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	content, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("error reading input file:", err)
		os.Exit(1)
	}

	blocks := strings.Split(string(content), "\n")
	towles, designs, err := parseInput(blocks)
	if err != nil {
		fmt.Println("parsing error:", err)
		os.Exit(1)
	}

	part1(towles, designs)
	part2(towles, designs)
}
