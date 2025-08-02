package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

func extractValues(match string) (int, int, error) {
	var length, repetitions int
	_, err := fmt.Sscanf(match, "(%dx%d)", &length, &repetitions)
	if err != nil {
		return 0, 0, nil
	}
	return length, repetitions, nil
}

func part1(line string) {
	count := 0
	idx := 0

	re := regexp.MustCompile(`\((\d+)x(\d+)\)`)
	matches := re.FindAllIndex([]byte(line), -1)

	for _, match := range matches {
		length, times, err := extractValues(line[match[0]:match[1]])
		if err != nil {
			continue
		}

		if match[0] < idx {
			continue
		}
		count += min(length, len(line)-idx)*times + (match[0] - idx)
		idx = match[1] + length

	}

	count += len(line) - idx
	fmt.Printf("Part 1: %d\n", count)
}

func part2(line string, first bool) int {
	count := 0
	idx := 0

	re := regexp.MustCompile(`\((\d+)x(\d+)\)`)
	matchesIdxs := re.FindAllIndex([]byte(line), -1)

	for _, match := range matchesIdxs {
		length, times, err := extractValues(line[match[0]:match[1]])
		if err != nil {
			continue
		}

		if match[0] < idx {
			continue
		}
		currLen := part2(line[match[1]:match[1]+length], false)
		count += currLen*times + (match[0] - idx)
		idx = match[1] + length

	}

	count += len(line) - idx
	if first {
		fmt.Println("Part 2:", count)
	}
	return count
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
	line := string(fileContent)

	part1(line)
	part2(line, true)
}
