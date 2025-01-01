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

func parseRotation(aRotation string) (dirMult, numClicks int, err error) {
	dirMult = -1
	if aRotation[0] == 'R' {
		dirMult = 1
	}

	numClicks, err = strconv.Atoi(aRotation[1:])
	return
}

func part1(rotations []string) {
	count := 0
	currPos := startingPosition

	for _, aRotation := range rotations {
		dirMult, numClicks, err := parseRotation(aRotation)
		if err != nil {
			continue
		}

		currPos = (currPos + (numClicks * dirMult) + dialSize) % dialSize
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
		dirMult, numClicks, err := parseRotation(aRotation)
		if err != nil {
			continue
		}

		count += numClicks / dialSize
		numClicks = numClicks % dialSize

		nextPos := numClicks*dirMult + currPos
		if currPos != 0 && (nextPos <= 0 || nextPos >= dialSize) {
			count++
		}

		currPos = (currPos + numClicks*dirMult + dialSize) % dialSize

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
