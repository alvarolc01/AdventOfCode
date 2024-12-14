package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func readFile(fileName *string) []string {
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

	return strings.Split(string(content), "\n")

}

func part1(input []string) {
	totalPaper := 0
	var height int
	var width int
	var length int
	for _, line := range input {
		_, err := fmt.Sscanf(line, "%dx%dx%d", &length, &width, &height)
		if err != nil {
			fmt.Printf("Failed to parse line: %s\n", line)
			continue
		}

		paperPresent := 2*length*width + 2*width*height + 2*height*length
		paperPresent += min(length*width, width*height, height*length)

		totalPaper += paperPresent
	}

	fmt.Printf("Part1: %d\n", totalPaper)
}

func part2(input []string) {
	totalPaper := 0
	var height int
	var width int
	var length int
	for _, line := range input {
		_, err := fmt.Sscanf(line, "%dx%dx%d", &length, &width, &height)
		if err != nil {
			fmt.Printf("Failed to parse line: %s\n", line)
			continue
		}

		paperPresent := length * width * height

		largestSide := max(length, width, height)
		paperPresent += 2*length + 2*width + 2*height - 2*largestSide

		totalPaper += paperPresent
	}

	fmt.Printf("Part1: %d\n", totalPaper)

}

func main() {
	fileName := flag.String("file", "", "Path to the file to read")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("File name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	inputStrings := readFile(fileName)

	part1(inputStrings)
	part2(inputStrings)

}
