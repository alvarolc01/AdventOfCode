package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	startingPosition int = 50
	dialSize         int = 100
)

func part1(rotations []string) {
	count := 0
	currPos := startingPosition

	for _, aRotation := range rotations {
		mult := -1
		if aRotation[0] == 'R' {
			mult = 1
		}

		numPositions, err := strconv.Atoi(aRotation[1:])
		if err != nil {
			continue
		}

		currPos = (currPos + (numPositions * mult) + dialSize) % dialSize
		if currPos == 0 {
			count++
		}

	}

	fmt.Printf("Part 1: %d\n", count)
}

func part2(rotations []string) {
	count := 0
	currPos := startingPosition

	for _, aRotation := range rotations {
		mult := -1
		if aRotation[0] == 'R' {
			mult = 1
		}

		numPositions, err := strconv.Atoi(aRotation[1:])
		if err != nil {
			continue
		}

		count += numPositions / dialSize
		numPositions = numPositions % dialSize

		if currPos != 0 && (numPositions*mult+currPos <= 0 || numPositions*mult+currPos >= dialSize) {
			count++
		}

		currPos = (currPos + numPositions*mult + dialSize) % dialSize

	}

	fmt.Printf("Part 2: %d\n", count)
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
