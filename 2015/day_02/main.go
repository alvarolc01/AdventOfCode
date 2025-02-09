package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type BoxDimensions struct {
	Width  int
	Height int
	Length int
}

func NewBoxDimensions(line string) *BoxDimensions {
	var height, width, length int
	_, err := fmt.Sscanf(line, "%dx%dx%d", &length, &width, &height)
	if err != nil {
		return nil
	}
	return &BoxDimensions{
		Width:  width,
		Height: height,
		Length: length,
	}
}

func part1(input []string) {
	totalPaper := 0

	for _, line := range input {
		box := NewBoxDimensions(line)
		if box == nil {
			continue
		}

		firstSide := box.Length * box.Width
		secondSide := box.Width * box.Height
		thirdSide := box.Height * box.Length

		paperPresent := 2*firstSide + 2*secondSide + 2*thirdSide
		paperPresent += min(firstSide, secondSide, thirdSide)

		totalPaper += paperPresent
	}

	fmt.Printf("Part 1: %d\n", totalPaper)
}

func part2(input []string) {
	totalPaper := 0

	for _, line := range input {
		box := NewBoxDimensions(line)
		if box == nil {
			continue
		}

		paperPresent := box.Length * box.Width * box.Height

		largestSide := max(box.Length, box.Width, box.Height)
		paperPresent += 2*box.Length + 2*box.Width + 2*box.Height - 2*largestSide

		totalPaper += paperPresent
	}

	fmt.Printf("Part 2: %d\n", totalPaper)
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
		fmt.Printf("Error reading input file: %s\n", err)
		os.Exit(1)
	}
	input := strings.Split(string(fileContent), "\n")

	part1(input)
	part2(input)
}
