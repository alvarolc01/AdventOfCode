package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(input []string) {
	sumDifference := 0

	for _, line := range input {
		unquotedLine, err := strconv.Unquote(line)
		if err != nil {
			fmt.Println("Failed to unquote line", line)
		} else {
			sumDifference += len(line) - len(unquotedLine)
		}
	}

	fmt.Printf("Part 1: %d\n", sumDifference)
}

func part2(input []string) {
	sumDifference := 0

	for _, line := range input {
		quotedLine := strconv.Quote(line)
		sumDifference += len(quotedLine) - len(line)
	}

	fmt.Printf("Part 2: %d\n", sumDifference)
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
