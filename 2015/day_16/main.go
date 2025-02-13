package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func NewAunt(line string) map[string]int {
	var auntId, val1, val2, val3 int
	var field1, field2, field3 string
	numFields, err := fmt.Sscanf(line, "Sue %d: %s %d, %s %d, %s %d", &auntId, &field1, &val1, &field2, &val2, &field3, &val3)
	if err != nil || numFields != 7 {
		return nil
	}

	return map[string]int{
		strings.TrimSuffix(field1, ":"): val1,
		strings.TrimSuffix(field2, ":"): val2,
		strings.TrimSuffix(field3, ":"): val3,
	}
}

func isAuntExact(aunt, knownAunt map[string]int) bool {
	for key, value := range aunt {
		if value != knownAunt[key] {
			return false
		}
	}
	return true
}

func part1(input []string, knownAunt map[string]int) {
	foundId := 0
	for idx := 0; idx < len(input) && foundId == 0; idx++ {
		currentAunt := NewAunt(input[idx])
		if isAuntExact(currentAunt, knownAunt) {
			foundId = idx + 1
		}
	}

	fmt.Printf("Part 1: %d\n", foundId)
}

func isAuntWithRanges(aunt, knownAunt map[string]int) bool {
	for key, value := range aunt {
		if key == "cats" || key == "trees" {
			if value <= knownAunt[key] {
				return false
			}
		} else if key == "pomeranians" || key == "goldfish" {
			if value >= knownAunt[key] {
				return false
			}
		} else {
			if value != knownAunt[key] {
				return false
			}
		}
	}
	return true
}

func part2(input []string, knownAunt map[string]int) {
	foundId := 0
	for idx := 0; idx < len(input) && foundId == 0; idx++ {
		currentAunt := NewAunt(input[idx])
		if isAuntWithRanges(currentAunt, knownAunt) {
			foundId = idx + 1
		}
	}

	fmt.Printf("Part 2: %d\n", foundId)
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

	knownAunt := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	part1(input, knownAunt)
	part2(input, knownAunt)
}
