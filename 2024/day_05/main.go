package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type PrecedenceMap map[int]map[int]bool

func isValidMagazine(magazine []int, precedences PrecedenceMap) bool {
	return slices.IsSortedFunc(magazine, func(a, b int) int {
		if precedences[b][a] {
			return 1
		}
		return -1
	})
}

func part1(precedences PrecedenceMap, magazines [][]int) {
	midPagesSum := 0

	for _, mag := range magazines {
		if isValidMagazine(mag, precedences) {
			midPagesSum += mag[len(mag)/2]
		}
	}

	fmt.Printf("Part 1: %d\n", midPagesSum)
}

func sortMagazineWithPrecedence(magazine []int, precedences PrecedenceMap) {
	sort.Slice(magazine, func(i, j int) bool {
		if precedences[magazine[i]][magazine[j]] {
			return true
		}
		if precedences[magazine[j]][magazine[i]] {
			return false
		}
		return true
	})
}

func part2(precedences PrecedenceMap, magazines [][]int) {
	midPagesSum := 0

	for _, mag := range magazines {
		if !isValidMagazine(mag, precedences) {
			sortMagazineWithPrecedence(mag, precedences)
			midPagesSum += mag[len(mag)/2]
		}
	}

	fmt.Printf("Part 2: %d\n", midPagesSum)
}

func parsePrecedences(section string) (PrecedenceMap, error) {
	precedences := make(PrecedenceMap)
	lines := strings.Split(strings.TrimSpace(section), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		var before, after int
		_, err := fmt.Sscanf(line, "%d|%d", &before, &after)
		if err != nil {
			return nil, fmt.Errorf("line (%q): invalid format", line)
		}

		if _, ok := precedences[before]; !ok {
			precedences[before] = make(map[int]bool)
		}
		precedences[before][after] = true
	}

	return precedences, nil
}

func parseMagazines(section string) ([][]int, error) {
	var magazines [][]int
	lines := strings.Split(strings.TrimSpace(section), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		numStrs := strings.Split(line, ",")
		numbers := make([]int, 0, len(numStrs))

		for _, numStr := range numStrs {
			numStr = strings.TrimSpace(numStr)
			n, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, fmt.Errorf("line (%q): invalid number %q", line, numStr)
			}
			numbers = append(numbers, n)
		}

		magazines = append(magazines, numbers)
	}

	return magazines, nil
}

func parseInput(input string) (PrecedenceMap, [][]int, error) {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	if len(sections) != 2 {
		return nil, nil, fmt.Errorf("expected two sections separated by a blank line")
	}

	precedences, err := parsePrecedences(sections[0])
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse precedences: %w", err)
	}

	magazines, err := parseMagazines(sections[1])
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse magazines: %w", err)
	}

	return precedences, magazines, nil
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
	precedences, magazines, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(precedences, magazines)
	part2(precedences, magazines)
}
