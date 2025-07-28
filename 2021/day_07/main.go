package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(nums []int) {
	maxElem := slices.Max(nums)
	countPerVal := make([]int, maxElem+1)
	for _, n := range nums {
		countPerVal[n]++
	}

	costToMeetAt := make([]int, len(countPerVal))
	foundSoFar := 0
	costSoFar := 0

	for i := 0; i < len(countPerVal); i++ {
		costToMeetAt[i] = costSoFar
		foundSoFar += countPerVal[i]
		costSoFar += foundSoFar
	}

	foundSoFar = 0
	costSoFar = 0

	for i := len(countPerVal) - 1; i >= 0; i-- {
		costToMeetAt[i] += costSoFar
		foundSoFar += countPerVal[i]
		costSoFar += foundSoFar
	}

	minCost := slices.Min(costToMeetAt)

	fmt.Printf("Part 1: %d\n", minCost)
}

func individualSideCosts(arr []int) []int {
	cost := make([]int, len(arr))

	totalCount := 0
	totalI := 0
	totalI2 := 0

	for i := range len(arr) {
		lc := (i*i*totalCount - 2*i*totalI + totalI2 + i*totalCount - totalI) / 2
		cost[i] = lc

		totalCount += arr[i]
		totalI += i * arr[i]
		totalI2 += i * i * arr[i]
	}

	return cost
}

func updatedCostsAtPosition(arr []int) []int {
	leftCost := individualSideCosts(arr)

	slices.Reverse(arr)
	rightCost := individualSideCosts(arr)
	slices.Reverse(rightCost)

	totalCosts := make([]int, len(arr))
	for i := range len(arr) {
		totalCosts[i] = leftCost[i] + rightCost[i]
	}

	return totalCosts
}

func part2(nums []int) {
	maxElem := slices.Max(nums)
	countPerVal := make([]int, maxElem+1)
	for _, n := range nums {
		countPerVal[n]++
	}

	ans := updatedCostsAtPosition(countPerVal)
	fmt.Printf("Part 2: %d\n", slices.Min(ans))
}

func parseInput(input string) ([]int, error) {
	nums := strings.Split(input, ",")
	output := make([]int, len(nums))
	for idx, line := range nums {
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
	input := string(fileContent)

	nums, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(nums)
	part2(nums)
}
