package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func containsThreeVowels(line string) bool {
	countVowels := 0
	for _, char := range line {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			countVowels++
		}
	}
	return countVowels >= 3
}

func containsRepeatedChars(line string) bool {
	for idx := 1; idx < len(line); idx++ {
		if line[idx] == line[idx-1] {
			return true
		}
	}
	return false
}

func containsForbiddenSubstrings(line string) bool {
	pattern := regexp.MustCompile("(ab|cd|pq|xy)")
	return pattern.MatchString(line)
}

func part1(input []string) {
	niceStrings := 0
	for _, line := range input {
		if containsThreeVowels(line) && containsRepeatedChars(line) && !containsForbiddenSubstrings(line) {
			niceStrings++
		}
	}

	fmt.Printf("Part 1: %d\n", niceStrings)
}

func hasPalindromeSubstring(line string) bool {
	for idx := 2; idx < len(line); idx++ {
		if line[idx-2] == line[idx] {
			return true
		}
	}
	return false
}

func hasNonOverlappingMatch(line string) bool {
	foundPairs := make(map[string]int)

	for idx := 1; idx < len(line); idx++ {
		currSubstr := line[idx-1 : idx+1]
		firstFoundIdx, found := foundPairs[currSubstr]
		if !found {
			foundPairs[currSubstr] = idx
		} else if firstFoundIdx != idx-1 {
			return true
		}
	}
	return false
}

func part2(input []string) {
	niceStrings := 0
	for _, line := range input {
		if hasNonOverlappingMatch(line) && hasPalindromeSubstring(line) {
			niceStrings++
		}
	}

	fmt.Printf("Part 2: %d\n", niceStrings)
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

	part1(input)
	part2(input)
}
