package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	rows = 6
	cols = 50
)

type Grid [][]bool

func NewGrid() Grid {
	g := make([][]bool, rows)
	for row := range g {
		g[row] = make([]bool, cols)
	}

	return g
}

func (g Grid) ApplyInstruction(line string) {
	switch {
	case strings.HasPrefix(line, "rect"):
		g.rect(line)
	case strings.HasPrefix(line, "rotate row"):
		g.rotateRow(line)
	case strings.HasPrefix(line, "rotate column"):
		g.rotateCol(line)
	default:
		fmt.Println("unknown instruction:", line)
	}
}

func (g Grid) rect(line string) {
	var width, height int
	_, err := fmt.Sscanf(line, "rect %dx%d", &width, &height)
	if err != nil {
		fmt.Println("err rect", line)
		return
	}
	for i := range height {
		for j := range width {
			g[i][j] = true
		}
	}
}

func (g Grid) rotateCol(line string) {
	var col, times int
	_, err := fmt.Sscanf(line, "rotate column x=%d by %d", &col, &times)
	if err != nil {
		fmt.Println("err col", line)
		return
	}

	nextCol := make([]bool, rows)
	for i := range rows {
		nextCol[(i+times)%rows] = g[i][col]
	}
	for i := range rows {
		g[i][col] = nextCol[i]
	}
}

func (g Grid) rotateRow(line string) {
	var row, times int
	_, err := fmt.Sscanf(line, "rotate row y=%d by %d", &row, &times)
	if err != nil {
		fmt.Println("err row")
		return
	}

	nextRow := make([]bool, cols)
	for i := range cols {
		nextRow[(i+times)%cols] = g[row][i]
	}
	g[row] = nextRow
}

func part1(commands []string) {
	grid := NewGrid()

	for _, command := range commands {
		grid.ApplyInstruction(command)
	}

	count := 0
	for _, i := range grid {
		for _, j := range i {
			if j {
				count++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", count)

}

func part2(commands []string) {
	grid := NewGrid()

	for _, command := range commands {
		grid.ApplyInstruction(command)
	}

	fmt.Println("Part 2:")

	for _, i := range grid {
		for _, j := range i {
			if j {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
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
	commands := strings.Split(string(fileContent), "\n")

	part1(commands)
	part2(commands)
}
