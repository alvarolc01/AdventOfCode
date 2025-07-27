package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	Part1WindowSize = 4
	Part2WindowSize = 14
)

func findMarker(signal string, windowSize int) int {
	lastChars := make(map[byte]int)
	pos := 0

	for ; pos < len(signal) && len(lastChars) != windowSize; pos++ {
		if pos >= windowSize {
			outisdeChar := signal[pos-windowSize]
			repetitions := lastChars[outisdeChar]
			if repetitions == 1 {
				delete(lastChars, outisdeChar)
			} else {
				lastChars[outisdeChar]--
			}
		}
		lastChars[signal[pos]]++
	}

	return pos
}

func part1(signal string) {
	posMarker := findMarker(signal, Part1WindowSize)
	fmt.Printf("Part 1: %d\n", posMarker)
}

func part2(signal string) {
	posMarker := findMarker(signal, Part2WindowSize)
	fmt.Printf("Part 2: %d\n", posMarker)
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

	part1(input)
	part2(input)
}
