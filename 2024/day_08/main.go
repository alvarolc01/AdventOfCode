package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	row, col int
}

type Grid struct {
	antenas       map[rune][]Point
	height, width int
}

func (g *Grid) isWithinBounds(p Point) bool {
	return p.row >= 0 && p.col >= 0 && p.row < g.height && p.col < g.width
}

func part1(grid *Grid) {
	antinodes := make(map[Point]bool)
	for _, antennas := range grid.antenas {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				dx := antennas[i].row - antennas[j].row
				dy := antennas[i].col - antennas[j].col

				p1 := Point{row: antennas[i].row + dx, col: antennas[i].col + dy}
				if grid.isWithinBounds(p1) {
					antinodes[p1] = true
				}

				p2 := Point{row: antennas[j].row - dx, col: antennas[j].col - dy}
				if grid.isWithinBounds(p2) {
					antinodes[p2] = true
				}

			}
		}
	}

	fmt.Printf("Part 1: %d\n", len(antinodes))
}

func addAntinodesInLine(start Point, dx, dy int, g *Grid, antinodes map[Point]bool) {
	p := start
	for g.isWithinBounds(p) {
		antinodes[p] = true
		p.row += dx
		p.col += dy
	}
}

func part2(grid *Grid) {
	antinodes := make(map[Point]bool)
	for _, antennas := range grid.antenas {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				dx := antennas[i].row - antennas[j].row
				dy := antennas[i].col - antennas[j].col

				addAntinodesInLine(antennas[i], dx, dy, grid, antinodes)
				addAntinodesInLine(antennas[j], -dx, -dy, grid, antinodes)
			}
		}
	}

	fmt.Printf("Part 2: %d\n", len(antinodes))
}

func parseInput(lines []string) *Grid {
	result := &Grid{
		height:  len(lines),
		width:   len(lines[0]),
		antenas: make(map[rune][]Point),
	}

	for i, line := range lines {
		for j, char := range line {
			if char == '.' {
				continue
			}
			result.antenas[char] = append(result.antenas[char], Point{row: i, col: j})
		}
	}

	return result
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
	grid := parseInput(rows)

	part1(grid)
	part2(grid)
}
