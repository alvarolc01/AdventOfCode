package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

func reactedPolymerLength(polymer string) int {
	st := []rune{}

	for _, letter := range polymer {
		if len(st) > 0 && st[len(st)-1] != letter && (unicode.ToLower(letter) == unicode.ToLower(st[len(st)-1])) {
			st = st[:len(st)-1]
		} else {
			st = append(st, letter)
		}
	}

	return len(st)
}

func part1(polymer string) {
	reactedLength := reactedPolymerLength(polymer)

	fmt.Printf("Part 1: %d\n", reactedLength)
}

func part2(polymer string) {
	minLen := math.MaxInt

	for i := 'a'; i <= 'z'; i++ {
		modified := strings.Replace(polymer, string(i), "", -1)
		modified = strings.Replace(modified, string(unicode.ToUpper(i)), "", -1)

		reactedLength := reactedPolymerLength(modified)
		minLen = min(minLen, reactedLength)
	}

	fmt.Printf("Part 2: %d\n", minLen)
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
