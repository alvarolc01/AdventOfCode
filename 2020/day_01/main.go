package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const TARGET_NUM int = 2020

func part1(entries []int) {
	left := 0
	right := len(entries) - 1

	for left < right && entries[left]+entries[right] != TARGET_NUM {
		sum := entries[left] + entries[right]
		if sum < TARGET_NUM {
			left++
		} else if sum > TARGET_NUM {
			right--
		}
	}

	fmt.Printf("Part 1: %d\n", entries[left]*entries[right])
}

func part2(entries []int) {
	var first, second, third int
	found := false

	for i := 0; i < len(entries)-2 && !found; i++ {
		left := i + 1
		right := len(entries) - 1
		target := TARGET_NUM - entries[i]
		for left < right {
			sum := entries[left] + entries[right]
			if sum == target {
				first, second, third = entries[i], entries[left], entries[right]
				found = true
				break
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}

	fmt.Printf("Part 2: %d\n", first*second*third)
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

	slices.Sort(output)
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

	entries, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(entries)
	part2(entries)
}
