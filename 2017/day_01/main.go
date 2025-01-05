package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func readFile(fileName *string) string {
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

	return string(content)

}

func sumEqualDigits(input string, distanceToMatchingDigit int) int {
	sumDigits := 0

	for i := 0; i < len(input); i++ {
		nextPosistion := (i + distanceToMatchingDigit) % len(input)
		if input[i] == input[nextPosistion] {
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

	inputString := readFile(fileName)

	part1(inputString)
	part2(inputString)

}
