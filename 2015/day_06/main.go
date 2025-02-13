package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	gridSideLength int    = 1000
	turnOnPrefix   string = "turn on"
	turnOffPrefix  string = "turn off"
	togglePrefix   string = "toggle"
)

type Row []int

type Grid struct {
	rows []Row
}

func NewGrid(numRows, numCols int) *Grid {
	rowsGrid := make([]Row, numRows)
	for idxRow := range rowsGrid {
		rowsGrid[idxRow] = make([]int, numCols)
	}

	return &Grid{
		rows: rowsGrid,
	}
}

func (g *Grid) CountLitLights() int {
	lightsOn := 0

	for _, row := range g.rows {
		for _, lightBrightness := range row {
			lightsOn += min(lightBrightness, 1)
		}
	}

	return lightsOn
}

func (g *Grid) GetTotalBrightness() int {
	totalBrightness := 0

	for _, row := range g.rows {
		for _, lightBrightness := range row {
			totalBrightness += lightBrightness
		}
	}

	return totalBrightness
}

type coordinatesCommands struct {
	initialX, initialY, finalX, finalY int
}

func (g *Grid) ParseCommand(command string) *coordinatesCommands {
	re := regexp.MustCompile(`(\d+),(\d+) through (\d+),(\d+)`)
	matches := re.FindStringSubmatch(command)
	if len(matches) != 5 {
		return nil
	}

	initialX, _ := strconv.Atoi(matches[1])
	initialY, _ := strconv.Atoi(matches[2])
	finalX, _ := strconv.Atoi(matches[3])
	finalY, _ := strconv.Atoi(matches[4])
	return &coordinatesCommands{
		initialX: initialX,
		initialY: initialY,
		finalX:   finalX,
		finalY:   finalY,
	}
}

func (g *Grid) ApplyCommand(rangeValues *coordinatesCommands, change func(int) int) {
	for idxX := rangeValues.initialX; idxX <= rangeValues.finalX; idxX++ {
		for idxY := rangeValues.initialY; idxY <= rangeValues.finalY; idxY++ {
			g.rows[idxX][idxY] = change(g.rows[idxX][idxY])
		}
	}
}

func turnOn(_ int) int            { return 1 }
func turnOff(_ int) int           { return 0 }
func toggle(initialValue int) int { return 1 - initialValue }

func part1(input []string) {
	grid := NewGrid(gridSideLength, gridSideLength)

	for _, command := range input {
		rangeValues := grid.ParseCommand(command)
		if strings.HasPrefix(command, turnOnPrefix) {
			grid.ApplyCommand(rangeValues, turnOn)
		} else if strings.HasPrefix(command, turnOffPrefix) {
			grid.ApplyCommand(rangeValues, turnOff)
		} else if strings.HasPrefix(command, togglePrefix) {
			grid.ApplyCommand(rangeValues, toggle)
		}
	}

	litLights := grid.CountLitLights()

	fmt.Printf("Part 1: %d\n", litLights)
}

func turnBrightnessOn(initialValue int) int  { return initialValue + 1 }
func turnBrightnessOff(initialValue int) int { return max(0, initialValue-1) }
func toggleBrightness(initialValue int) int  { return initialValue + 2 }

func part2(input []string) {
	grid := NewGrid(gridSideLength, gridSideLength)

	for _, command := range input {
		rangeValues := grid.ParseCommand(command)
		if strings.HasPrefix(command, turnOnPrefix) {
			grid.ApplyCommand(rangeValues, turnBrightnessOn)
		} else if strings.HasPrefix(command, turnOffPrefix) {
			grid.ApplyCommand(rangeValues, turnBrightnessOff)
		} else if strings.HasPrefix(command, togglePrefix) {
			grid.ApplyCommand(rangeValues, toggleBrightness)
		}
	}

	totalBrightness := grid.GetTotalBrightness()

	fmt.Printf("Part 2: %d\n", totalBrightness)
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
