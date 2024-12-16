package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
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

func containsThreeVowels(line string) bool {
	vowelsFound := 0
	for _, char := range line {
		if strings.ContainsRune("aeiou", char) {
			vowelsFound++
		}
		if vowelsFound == 3 {
			return true
		}
	}
	return false
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

func isNice(line string) bool {
	return containsThreeVowels(line) && containsRepeatedChars(line) && !containsForbiddenSubstrings(line)
}

func part1(input []string) {
	niceStrings := 0
	for _, line := range input {
		if isNice(line) {
			niceStrings++
		}
	}

	fmt.Printf("Part 1: %d\n", niceStrings)
}

func hasPalindromeSubstring(line string) bool {
	for idx := 0; idx < len(line)-2; idx++ {
		if line[idx] == line[idx+2] {
			return true
		}
	}
	return false
}

func hasNonOverlappingMatch(line string) bool {
	for idx := 0; idx < len(line)-1; idx++ {
		currSubstr := line[idx : idx+2]

		for secondIdx := idx + 2; secondIdx < len(line)-1; secondIdx++ {
			if line[secondIdx:secondIdx+2] == currSubstr {
				return true
			}
		}

	}

	return false
}

func secondIsNice(line string) bool {
	return hasNonOverlappingMatch(line) && hasPalindromeSubstring(line)
}

func part2(input []string) {
	niceStrings := 0
	for _, line := range input {
		if secondIsNice(line) {
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

	inputStrings := readFile(fileName)

	part1(inputStrings)
	part2(inputStrings)

}
