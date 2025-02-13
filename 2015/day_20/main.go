package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

const (
	PresentsGivenFirstPart  int = 10
	PresentsGivenSecondPart int = 11
	MaxHousesVisitedPerElf  int = 50
)

func part1(minPresents int) {
	lastHouse := minPresents / PresentsGivenFirstPart
	housePresents := make([]int, lastHouse+1)

	elf := 1
	for ; elf <= lastHouse; elf++ {
		for houseVisited := elf; houseVisited <= lastHouse; houseVisited += elf {
			housePresents[houseVisited] += elf * PresentsGivenFirstPart
		}
		if housePresents[elf] >= minPresents {
			break
		}
	}

	fmt.Printf("Part 1: %d\n", elf)
}

func part2(minPresents int) {
	lastHouse := minPresents / PresentsGivenSecondPart
	housePresents := make([]int, lastHouse+1)

	elf := 1
	for ; elf <= lastHouse; elf++ {
		numHousesVisited := 1
		for houseVisited := elf; houseVisited <= lastHouse && numHousesVisited <= MaxHousesVisitedPerElf; houseVisited += elf {

			housePresents[houseVisited] += elf * PresentsGivenSecondPart
			numHousesVisited++
		}

		if housePresents[elf] >= minPresents {
			break
		}
	}

	fmt.Printf("Part 2: %d\n", elf)
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

	minPresents, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Failed to parse input strings: %s\n", err)
		os.Exit(1)
	}

	part1(minPresents)
	part2(minPresents)
}
