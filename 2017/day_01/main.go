package main

import (
	"flag"
	"fmt"
	"os"
)

func sumEqualDigits(input string, distanceToMatchingDigit int) int {
	sumDigits := 0

	for i := 0; i < len(input); i++ {
		nextPosition := (i + distanceToMatchingDigit) % len(input)
		if input[i] == input[nextPosition] {
			sumDigits += int(input[i] - '0')
		}
	}

	return sumDigits
}

func part1(input string) {
	sumDigits := sumEqualDigits(input, 1)

	fmt.Printf("Part 1: %d\n", sumDigits)
}

func part2(input string) {
	sumDigits := sumEqualDigits(input, len(input)/2)

	fmt.Printf("Part 2: %d\n", sumDigits)
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
	input := string(fileContent)

	part1(input)
	part2(input)
}
