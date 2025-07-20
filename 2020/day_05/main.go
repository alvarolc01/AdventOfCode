package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

func binarySearch(specification string, left, right int) int {
	for _, char := range specification {
		mid := (right + left) / 2

		if char == 'F' || char == 'L' {
			right = mid
		} else {
			left = mid + 1
		}
		mid = (right + left) / 2
	}

	return left
}

func getSeatID(speficication string) int {
	if len(speficication) != 10 {
		return -1
	}

	rowSpeficication := speficication[:7]
	colSpecicication := speficication[7:]

	row := binarySearch(rowSpeficication, 0, 127)
	col := binarySearch(colSpecicication, 0, 7)
	return row*8 + col
}

func part1(boardingPass []string) {
	maxSeatID := 0

	for _, seat := range boardingPass {
		seatID := getSeatID(seat)
		maxSeatID = max(maxSeatID, seatID)
	}

	fmt.Printf("Part 1: %d\n", maxSeatID)
}

func part2(boardingPass []string) {
	foundIDs := make([]int, len(boardingPass))
	for idx, seat := range boardingPass {
		seatID := getSeatID(seat)
		foundIDs[idx] = seatID
	}

	sort.Ints(foundIDs)

	missingID := -1
	for i := 0; i < len(foundIDs)-1 && missingID == -1; i++ {
		if foundIDs[i]+1 != foundIDs[i+1] {
			missingID = foundIDs[i] + 1
		}
	}

	fmt.Printf("Part 2: %d\n", missingID)
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
