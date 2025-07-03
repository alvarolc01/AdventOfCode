package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const GEAR_SYMBOL rune = '*'

var directions = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

type Position struct {
	row, col int
}

func extractNumber(input []string, row, col int) (int, Position) {
	left := col
	right := col

	for right+1 < len(input[row]) && unicode.IsDigit(rune(input[row][right+1])) {
		right++
	}

	for left-1 >= 0 && unicode.IsDigit(rune(input[row][left-1])) {
		left--
	}

	strNum := input[row][left : right+1]
	num, err := strconv.Atoi(strNum)
	if err != nil {
		return 0, Position{-1, -1}
	}

	return num, Position{row, left}
}

func adjacentNumbers(input []string, row, col int) map[Position]int {
	positionedNums := make(map[Position]int)

	for _, dir := range directions {
		x, y := row+dir[0], col+dir[1]
		if x < 0 || y < 0 || x >= len(input) || y >= len(input[0]) {
			continue
		}

		if unicode.IsDigit(rune(input[x][y])) {
			num, pos := extractNumber(input, x, y)
			positionedNums[pos] = num
		}
	}

	return positionedNums
}

func part1(input []string) {
	sumParts := 0
	usedNums := make(map[Position]bool)

	for idxRow, row := range input {
		for idxCol, col := range row {
			if (unicode.IsSymbol(col) || unicode.IsPunct(col)) && col != '.' {
				surroudingNums := adjacentNumbers(input, idxRow, idxCol)

				for pos, num := range surroudingNums {
					if !usedNums[pos] {
						usedNums[pos] = true
						sumParts += num
					}
				}

			}
		}
	}

	fmt.Printf("Part 1: %d\n", sumParts)
}

func part2(input []string) {
	sumGearRatios := 0

	for idxRow, row := range input {
		for idxCol, col := range row {
			if col == GEAR_SYMBOL {
				surroudingNums := adjacentNumbers(input, idxRow, idxCol)

				if len(surroudingNums) == 2 {
					ratio := 1
					for _, val := range surroudingNums {
						ratio *= val
					}
					sumGearRatios += ratio
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", sumGearRatios)
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

	part1(input)
	part2(input)
}
