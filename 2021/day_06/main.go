package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	DaysPart1 = 80
	DaysPart2 = 256
)

type MemPoint struct {
	daysRemaining, fishVal int
}

func calculateFishesInDays(fishes []int, days int) int {
	fishesPerRemainingDays := make([]int, 9)
	for _, n := range fishes {
		fishesPerRemainingDays[n]++
	}

	for ; days > 0; days-- {
		fishesPerRemainingDays = append(fishesPerRemainingDays[1:], fishesPerRemainingDays[0])
		fishesPerRemainingDays[6] += fishesPerRemainingDays[8]
	}

	total := 0
	for _, n := range fishesPerRemainingDays {
		total += n
	}

	return total
}

func part1(fishes []int) {
	totalAns := calculateFishesInDays(fishes, DaysPart1)

	fmt.Printf("Part 1: %d\n", totalAns)
}

func part2(fishes []int) {
	totalAns := calculateFishesInDays(fishes, DaysPart2)

	fmt.Printf("Part 2: %d\n", totalAns)
}

func parseInput(fishes string) ([]int, error) {
	fishValues := strings.Split(fishes, ",")
	output := make([]int, len(fishValues))
	for idx, val := range fishValues {
		parsedLine, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		output[idx] = parsedLine
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
	input := string(fileContent)

	lines, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(lines)
	part2(lines)
}
