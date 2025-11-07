package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(scanners map[int]int) {
	severity := 0

	for idx, depth := range scanners {
		period := 2 * (depth - 1)
		if (idx % period) == 0 {
			severity += idx * depth
		}
	}

	fmt.Printf("Part 1: %d\n", severity)
}

func part2(scanners map[int]int) {
	for delay := 0; ; delay++ {
		caught := false
		for idx, depth := range scanners {
			period := 2 * (depth - 1)
			if (idx+delay)%period == 0 {
				caught = true
				break
			}
		}
		if !caught {
			fmt.Printf("Part 2: %d\n", delay)
			return
		}
	}
}

func parseInput(lines []string) (map[int]int, error) {
	output := make(map[int]int)

	for _, aLine := range lines {
		parts := strings.Split(aLine, ":")
		idx, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, err
		}

		depth, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, err
		}

		output[idx] = depth
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

	scanners, err := parseInput(input)
	if err != nil {
		fmt.Println("error parting input:", err)
		os.Exit(1)
	}

	part1(scanners)
	part2(scanners)
}
