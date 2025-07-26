package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	target  int
	numbers []int
}

func NewEquation(line string) (*Equation, error) {
	parts := strings.Split(line, ":")
	target, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	numbers := strings.Fields(strings.TrimSpace(parts[1]))
	numbersEquation := make([]int, len(numbers))

	for idxNum, num := range numbers {
		numbersEquation[idxNum], err = strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
	}

	return &Equation{
		target:  target,
		numbers: numbersEquation,
	}, nil
}

func reachPart1(nums []int, target, currValue, pos int) bool {
	if pos == len(nums) {
		return currValue == target
	} else if currValue > target {
		return false
	}

	validWithSum := reachPart1(nums, target, currValue+nums[pos], pos+1)
	if validWithSum {
		return true
	}

	return reachPart1(nums, target, currValue*nums[pos], pos+1)
}

func part1(equations []*Equation) {
	sumValidTargets := 0

	for _, eq := range equations {
		if reachPart1(eq.numbers, eq.target, 0, 0) {
			sumValidTargets += eq.target
		}
	}

	fmt.Printf("Part 1: %d\n", sumValidTargets)
}

func reachPart2(nums []int, target, currValue, pos int) bool {
	if pos == len(nums) {
		return currValue == target
	} else if currValue > target {
		return false
	}

	if reachPart2(nums, target, currValue+nums[pos], pos+1) ||
		reachPart2(nums, target, currValue*nums[pos], pos+1) {
		return true
	}
	combinedNums := strconv.Itoa(currValue) + strconv.Itoa(nums[pos])
	nextNum, err := strconv.Atoi(combinedNums)
	if err != nil {
		return false
	}
	return reachPart2(nums, target, nextNum, pos+1)
}

func part2(equations []*Equation) {
	sumValidTargets := 0

	for _, eq := range equations {
		if reachPart2(eq.numbers, eq.target, 0, 0) {
			sumValidTargets += eq.target
		}
	}

	fmt.Printf("Part 2: %d\n", sumValidTargets)
}

func parseInput(rows []string) ([]*Equation, error) {
	equations := make([]*Equation, len(rows))

	for idxRow, row := range rows {
		equation, err := NewEquation(row)
		if err != nil {
			return nil, err
		}
		equations[idxRow] = equation
	}

	return equations, nil
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
	rows := strings.Split(string(fileContent), "\n")
	equations, err := parseInput(rows)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(equations)
	part2(equations)
}
