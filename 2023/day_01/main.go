package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var wordsToNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getCalibrationValues(line string) int {
	first := -1
	second := 0

	for _, char := range line {
		if unicode.IsNumber(char) {
			digit := int(char - '0')
			second = digit
			if first == -1 {
				first = digit
			}
		}
	}

	return first*10 + second
}

func part1(input []string) {
	sumCalibrationValues := 0

	for _, row := range input {
		sumCalibrationValues += getCalibrationValues(row)
	}

	fmt.Printf("Part 1: %d\n", sumCalibrationValues)
}

func extractNumberWord(line string, pos int) (int, bool) {
	for key, val := range wordsToNumbers {
		if strings.HasPrefix(line[pos:], key) {
			return val, true
		}
	}

	return 0, false
}

func getCalibrationValuesWithWords(line string) int {
	first := -1
	second := 0

	for idx, char := range line {
		currentDigit := -1
		if unicode.IsNumber(char) {
			currentDigit = int(char - '0')
		} else {
			number, found := extractNumberWord(line, idx)
			if found {
				currentDigit = number
			}
		}

		if currentDigit != -1 {
			second = currentDigit
			if first == -1 {
				first = currentDigit
			}
		}
	}

	return first*10 + second
}

func part2(input []string) {
	sumCalibrationValues := 0
	for _, row := range input {
		sumCalibrationValues += getCalibrationValuesWithWords(row)
	}
	fmt.Printf("Part 2: %d\n", sumCalibrationValues)
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
	input := strings.Split(string(fileContent), "\n")

	part1(input)
	part2(input)
}
