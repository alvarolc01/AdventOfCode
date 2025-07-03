package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Elf []int

func NewElf(block string) (Elf, error) {
	nums := strings.Split(block, "\n")

	caloriesElf := make(Elf, len(nums))
	for idx, line := range nums {
		convertedN, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		caloriesElf[idx] = convertedN
	}

	return caloriesElf, nil
}

func totalCalories(e Elf) int {
	sumCal := 0

	for _, calories := range e {
		sumCal += calories
	}

	return sumCal
}

func part1(elves []Elf) {
	maxCal := 0

	for _, elf := range elves {
		if caloriesElf := totalCalories(elf); maxCal < caloriesElf {
			maxCal = caloriesElf
		}
	}

	fmt.Printf("Part 1: %d\n", maxCal)
}

func part2(elves []Elf) {
	first, second, third := 0, 0, 0

	for _, elf := range elves {
		caloriesElf := totalCalories(elf)
		if caloriesElf >= first {
			first, second, third = caloriesElf, first, second
		} else if caloriesElf >= second {
			second, third = caloriesElf, second
		} else if caloriesElf > third {
			third = caloriesElf
		}
	}

	caloriesTopThree := first + second + third

	fmt.Printf("Part 2: %d\n", caloriesTopThree)

}

func parseInput(blocks []string) ([]Elf, error) {
	elves := make([]Elf, len(blocks))
	for idx, block := range blocks {
		elf, err := NewElf(block)
		if err != nil {
			return nil, err
		}

		elves[idx] = elf
	}

	return elves, nil
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
	input := strings.Split(string(fileContent), "\n\n")

	elves, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(elves)
	part2(elves)
}
