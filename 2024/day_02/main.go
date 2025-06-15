package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isReportSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	isIncreasing := report[0] < report[1]

	for idx := 1; idx < len(report); idx++ {
		reportIncreasing := report[idx] > report[idx-1]
		variation := int(math.Abs(float64(report[idx] - report[idx-1])))

		if reportIncreasing != isIncreasing || variation < 1 || variation > 3 {
			return false
		}
	}

	return true
}

func isReportSafeWithTolerance(report []int) bool {
	for idx := 0; idx < len(report); idx++ {

		updatedSlice := append([]int{}, report[:idx]...)
		updatedSlice = append(updatedSlice, report[idx+1:]...)
		if isReportSafe(updatedSlice) {
			return true
		}
	}

	return false
}

func part1(reports [][]int) {
	safeReports := 0
	for _, currReport := range reports {
		if isReportSafe(currReport) {
			safeReports++
		}
	}

	fmt.Printf("Part 1: %d\n", safeReports)
}

func part2(reports [][]int) {
	safeReports := 0
	for _, currReport := range reports {
		if isReportSafe(currReport) || isReportSafeWithTolerance(currReport) {
			safeReports++
		}
	}

	fmt.Printf("Part 2: %d\n", safeReports)
}

func parseInput(input string) ([][]int, error) {
	lines := strings.Split(input, "\n")
	rows := make([][]int, len(lines))

	for idx, line := range lines {

		numbersLine := strings.Split(line, " ")
		rows[idx] = make([]int, len(numbersLine))

		for idxNum, num := range numbersLine {
			convertedNum, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}

			rows[idx][idxNum] = convertedNum
		}

	}

	return rows, nil
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

	rows, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(rows)
	part2(rows)
}
