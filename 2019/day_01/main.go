package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getFuel(mass int) int {
	return int(mass/3) - 2
}

func part1(masses []int) {
	totalFuel := 0

	for _, mass := range masses {
		totalFuel += getFuel(mass)
	}

	fmt.Printf("Part 1: %d\n", totalFuel)
}

func part2(masses []int) {
	totalFuel := 0

	for _, mass := range masses {
		requiredFuel := getFuel(mass)
		for requiredFuel > 0 {
			totalFuel += requiredFuel
			requiredFuel = getFuel(requiredFuel)
		}
	}

	fmt.Printf("Part 2: %d\n", totalFuel)
}

func parseInput(lines []string) ([]int, error) {
	output := make([]int, len(lines))

	for idx, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		output[idx] = num
	}

	return output, nil
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

	masses, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(masses)
	part2(masses)
}
