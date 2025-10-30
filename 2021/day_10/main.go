package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

var openingChars = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var errorScore = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func getSyntaxErrorScore(line string) int {
	st := []rune{}
	for _, ch := range line {
		if ch == '{' || ch == '[' || ch == '<' || ch == '(' {
			st = append(st, ch)
		} else if len(st) > 0 {
			expectedOpening := openingChars[ch]
			if st[len(st)-1] != expectedOpening {
				return errorScore[ch]
			}
			st = st[:len(st)-1]
		}
	}
	return 0
}

var closingScores = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func getCompletionScore(line string) int {
	st := []rune{}
	for _, ch := range line {
		if ch == '{' || ch == '[' || ch == '<' || ch == '(' {
			st = append(st, ch)
		} else if len(st) > 0 {
			expectedOpening := openingChars[ch]
			if st[len(st)-1] == expectedOpening {
				st = st[:len(st)-1]
			}
		}
	}
	completionScore := 0

	for len(st) > 0 {
		curr := st[len(st)-1]
		st = st[:len(st)-1]
		completionScore = completionScore*5 + closingScores[curr]
	}

	return completionScore
}

func part1(lines []string) {
	totalSyntaxErrorScore := 0

	for _, line := range lines {
		totalSyntaxErrorScore += getSyntaxErrorScore(line)
	}

	fmt.Printf("Part 1: %d\n", totalSyntaxErrorScore)
}

func part2(lines []string) {
	closingScores := []int{}

	for _, line := range lines {
		if getSyntaxErrorScore(line) == 0 {
			syntaxErrorScore := getCompletionScore(line)
			if syntaxErrorScore != 0 {
				closingScores = append(closingScores, syntaxErrorScore)
			}
		}
	}
	sort.Ints(closingScores)

	fmt.Printf("Part 1: %d\n", closingScores[len(closingScores)/2])
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

	part1(lines)
	part2(lines)
}
