package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const targetWord = "XMAS"

var directions = [][]int{
	{1, 1}, {1, 0}, {1, -1}, {0, -1},
	{-1, -1}, {-1, 0}, {-1, 1}, {0, 1},
}

func matchesWord(matrix []string, dir []int, row, col, idx int) bool {
	if idx == len(targetWord) {
		return true
	}

	if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[0]) {
		return false
	}

	if targetWord[idx] != matrix[row][col] {
		return false
	}

	return matchesWord(matrix, dir, row+dir[0], col+dir[1], idx+1)
}

func countWordsAround(matrix []string, row, col int) int {
	wordsAround := 0
	for _, dir := range directions {
		if matchesWord(matrix, dir, row, col, 0) {
			wordsAround++
		}
	}

	return wordsAround
}

func part1(matrix []string) {
	countApparitions := 0

	for idxRow, row := range matrix {
		for idxCol, letter := range row {
			if letter == 'X' {
				countApparitions += countWordsAround(matrix, idxRow, idxCol)
			}
		}
	}

	fmt.Printf("Part 1: %d\n", countApparitions)
}

func isMS(ch byte) bool {
	return ch == 'M' || ch == 'S'
}

func isCrossedMax(matrix []string, row, col int) bool {
	topLeft, topRight := matrix[row-1][col-1], matrix[row-1][col+1]
	bottomLeft, bottomRight := matrix[row+1][col-1], matrix[row+1][col+1]

	vertical := bottomLeft == topLeft && isMS(bottomLeft) && bottomRight == topRight && isMS(bottomRight) && bottomLeft != bottomRight
	horizontal := topRight == topLeft && isMS(topRight) && bottomRight == bottomLeft && isMS(bottomLeft) && bottomLeft != topLeft

	return vertical || horizontal
}

func part2(matrix []string) {
	countApparitions := 0

	for idxRow := 1; idxRow < len(matrix)-1; idxRow++ {
		for idxCol := 1; idxCol < len(matrix[0])-1; idxCol++ {
			if matrix[idxRow][idxCol] == 'A' && isCrossedMax(matrix, idxRow, idxCol) {
				countApparitions++
			}
		}
	}

	fmt.Printf("Part 2: %d\n", countApparitions)
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
