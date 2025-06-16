package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	enablingInstruction  = "do()"
	disablingInstruction = "don't()"
)

var (
	reBase     = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	reEnabling = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`)
)

func extractOperands(match string) (int, int, error) {
	var a, b int
	_, err := fmt.Sscanf(match, "mul(%d,%d)", &a, &b)
	return a, b, err
}

func part1(rows []string) {
	totalMult := 0

	for _, line := range rows {
		matches := reBase.FindAllString(line, -1)
		for _, match := range matches {
			multiplicand, multiplier, err := extractOperands(match)
			if err != nil {
				return
			}
			totalMult += multiplicand * multiplier
		}
	}

	fmt.Printf("Part 1: %d\n", totalMult)
}

func part2(rows []string) {
	totalMult := 0
	enabled := true

	for _, line := range rows {
		matches := reEnabling.FindAllString(line, -1)
		for _, match := range matches {
			if match == disablingInstruction {
				enabled = false
			} else if match == enablingInstruction {
				enabled = true
			} else if enabled {
				multiplicand, multiplier, err := extractOperands(match)
				if err != nil {
					return
				}
				totalMult += multiplicand * multiplier
			}

		}

	}

	fmt.Printf("Part 2: %d\n", totalMult)
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

	part1(rows)
	part2(rows)
}
