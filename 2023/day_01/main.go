package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

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
}
