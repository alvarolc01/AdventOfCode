package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(measurements []int) {
	increasingMeasurements := 0

	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			increasingMeasurements++
		}
	}

	fmt.Printf("Part 1: %d\n", increasingMeasurements)
}

func part2(measurements []int) {
	increasingMeasurements := 0

	for i := 3; i < len(measurements); i++ {
		if measurements[i] > measurements[i-3] {
			increasingMeasurements++
		}
	}

	fmt.Printf("Part 2: %d\n", increasingMeasurements)

}

func parseInput(lines []string) ([]int, error) {
	output := make([]int, len(lines))
	for idx, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		output[idx] = n
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

	measurements, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(measurements)
	part2(measurements)
}
