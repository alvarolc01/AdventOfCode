package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	initialPositionSecondPart int = 0
	numPositionsSecondPart    int = 11
)

type Disc struct {
	initialPosition, numPositions int
}

func (d *Disc) PositionAt(t int) int {
	return (d.initialPosition + t) % d.numPositions
}

func NewDisc(line string) (*Disc, error) {
	var initialPosition, numPositions int
	var numDisc int
	_, err := fmt.Sscanf(line, "Disc #%d has %d positions; at time=0, it is at position %d.", &numDisc, &numPositions, &initialPosition)
	if err != nil {
		return nil, err
	}

	return &Disc{
		initialPosition: initialPosition,
		numPositions:    numPositions,
	}, nil
}

func isValidTime(attemptedTime int, discs []*Disc) bool {
	for _, disc := range discs {
		attemptedTime++

		if disc.PositionAt(attemptedTime) != 0 {
			return false
		}
	}
	return true
}

func findFirstValidTime(discs []*Disc) int {
	for currTime := 0; ; currTime++ {
		if isValidTime(currTime, discs) {
			return currTime
		}
	}
}

func part1(discs []*Disc) {
	fmt.Printf("Part 1: %d\n", findFirstValidTime(discs))
}

func part2(discs []*Disc) {
	discs = append(discs, &Disc{
		initialPosition: initialPositionSecondPart,
		numPositions:    numPositionsSecondPart,
	})

	fmt.Printf("Part 2: %d\n", findFirstValidTime(discs))
}

func parseInput(lines []string) ([]*Disc, error) {
	result := make([]*Disc, 0, len(lines))
	for _, line := range lines {
		currDisc, err := NewDisc(line)
		if err != nil {
			return nil, err
		}
		result = append(result, currDisc)
	}
	return result, nil
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
	lines := strings.Split(string(fileContent), "\n")
	discs, err := parseInput(lines)
	if err != nil {
		fmt.Println("error reading input file:", err)
		os.Exit(1)

	}
	part1(discs)
	part2(discs)
}
