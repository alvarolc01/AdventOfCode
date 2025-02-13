package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const NumSteps int = 100

type Row []rune

type Grid struct {
	rows []Row
}

func NewGrid(initialGrid []string) *Grid {
	rowsGrid := make([]Row, len(initialGrid))
	for idxRow, rowContent := range initialGrid {
		rowsGrid[idxRow] = Row(rowContent)
	}

	return &Grid{
		rows: rowsGrid,
	}
}

func (g *Grid) CountLitNeighbours(idxRow, idxCol int) int {
	litNeighbours := 0

	for row := idxRow - 1; row <= idxRow+1; row++ {
		for col := idxCol - 1; col <= idxCol+1; col++ {
			isCenterLight := row == idxRow && col == idxCol
			isWithinBounds := row >= 0 && col >= 0 && row < len(g.rows) && col < len(g.rows[0])
			if !isCenterLight && isWithinBounds {
				if g.rows[row][col] == '#' {
					litNeighbours++
				}
			}
		}
	}

	return litNeighbours
}

func (g *Grid) Step() {
	rows, cols := len(g.rows), len(g.rows[0])
	nextRows := make([]Row, rows)
	for i := range nextRows {
		nextRows[i] = make(Row, cols)
	}

	for idxRow := 0; idxRow < rows; idxRow++ {
		for idxCol := 0; idxCol < cols; idxCol++ {
			litNeighbours := g.CountLitNeighbours(idxRow, idxCol)
			if g.rows[idxRow][idxCol] == '#' && (litNeighbours != 2 && litNeighbours != 3) {
				nextRows[idxRow][idxCol] = '.'
			} else if g.rows[idxRow][idxCol] == '.' && litNeighbours == 3 {
				nextRows[idxRow][idxCol] = '#'
			} else {
				nextRows[idxRow][idxCol] = g.rows[idxRow][idxCol]
			}
		}
	}
	g.rows = nextRows
}

func (g *Grid) StepStuckLights() {
	g.Step()
	g.setCorners()
}

func (g *Grid) setCorners() {
	g.rows[0][0] = '#'
	g.rows[0][len(g.rows[0])-1] = '#'
	g.rows[len(g.rows)-1][0] = '#'
	g.rows[len(g.rows)-1][len(g.rows[0])-1] = '#'
}

func (g *Grid) CountLitLights() int {
	lightsOn := 0

	for _, row := range g.rows {
		lightsOn += strings.Count(string(row), "#")
	}

	return lightsOn
}

func part1(input []string) {
	grid := NewGrid(input)

	for i := 0; i < NumSteps; i++ {
		grid.Step()
	}

	litLights := grid.CountLitLights()

	fmt.Printf("Part 1: %d\n", litLights)
}

func part2(input []string) {
	grid := NewGrid(input)
	grid.setCorners()

	for i := 0; i < NumSteps; i++ {
		grid.StepStuckLights()
	}

	litLights := grid.CountLitLights()

	fmt.Printf("Part 2: %d\n", litLights)
}

func main() {
	fileName := flag.String("file", "", "Path to the file to read")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("File name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	fileContent, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}
	input := strings.Split(string(fileContent), "\n")

	part1(input)
	part2(input)
}
