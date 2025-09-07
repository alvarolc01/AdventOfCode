package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	safeTileMark rune = '.'
	trapMark     rune = '^'

	tileRowsFirstPart  int = 40
	tileRowsSecondPart int = 400000
)

var trapTileCombinations = map[string]bool{
	"^^.": true,
	".^^": true,
	"..^": true,
	"^..": true,
}

func getCurrentTile(previousTiles string) rune {
	if _, isNextTrap := trapTileCombinations[previousTiles]; isNextTrap {
		return trapMark
	}
	return safeTileMark
}

func findNextTileRow(currentRow string) string {
	result := make([]rune, 0, len(currentRow))

	expandedRow := string(safeTileMark) + currentRow + string(safeTileMark)
	for i := 1; i < len(expandedRow)-1; i++ {
		result = append(result, getCurrentTile(expandedRow[i-1:i+2]))
	}

	return string(result)
}

func safeTilesInNRows(initialTiles string, rows int) int {
	result := 0
	currentRow := initialTiles
	for i := 0; i < rows; i++ {
		result += strings.Count(currentRow, string(safeTileMark))
		currentRow = findNextTileRow(currentRow)
	}

	return result
}

func part1(tiles string) {
	fmt.Printf("Part 1: %d\n", safeTilesInNRows(tiles, tileRowsFirstPart))
}

func part2(tiles string) {
	fmt.Printf("Part 2: %d\n", safeTilesInNRows(tiles, tileRowsSecondPart))
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

	initialTiles := string(fileContent)

	part1(initialTiles)
	part2(initialTiles)
}
