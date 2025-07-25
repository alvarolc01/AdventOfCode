package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func parseCommand(command string) (*Point, int, error) {
	var dir rune
	var n int
	_, err := fmt.Sscanf(command, "%c%d", &dir, &n)
	if err != nil {
		return nil, 0, err
	}
	var step Point
	switch dir {
	case 'U':
		step = Point{0, 1}
	case 'D':
		step = Point{0, -1}
	case 'R':
		step = Point{1, 0}
	case 'L':
		step = Point{-1, 0}
	}

	return &step, n, nil
}

type Wire struct {
	coordinates map[Point]int
}

func NewWire(line string) (*Wire, error) {
	commands := strings.Split(line, ",")
	wireCoordinates := make(map[Point]int)
	currPosition := Point{0, 0}
	distance := 0
	for _, command := range commands {
		dir, steps, err := parseCommand(command)
		if err != nil {
			return nil, err
		}

		for i := 0; i < steps; i++ {
			currPosition.x += dir.x
			currPosition.y += dir.y
			distance++
			if _, ok := wireCoordinates[currPosition]; !ok {
				wireCoordinates[currPosition] = distance
			}
		}
	}
	return &Wire{
		coordinates: wireCoordinates,
	}, nil
}

type Panel struct {
	first, second *Wire
}

func NewPanel(lines []string) (*Panel, error) {
	if len(lines) != 2 {
		return nil, fmt.Errorf("required 2 wires")
	}

	firstWire, err := NewWire(lines[0])
	if err != nil {
		return nil, err
	}

	secondWire, err := NewWire(lines[1])
	if err != nil {
		return nil, err
	}

	return &Panel{
		first:  firstWire,
		second: secondWire,
	}, nil

}

func part1(panel *Panel) {
	distanceClosestIntersection := math.MaxInt

	for pointFirstWire := range panel.first.coordinates {
		if _, ok := panel.second.coordinates[pointFirstWire]; ok {
			distance := int(math.Abs(float64(pointFirstWire.x)) + math.Abs(float64(pointFirstWire.y)))
			distanceClosestIntersection = min(distanceClosestIntersection, distance)
		}
	}

	fmt.Printf("Part 1: %d\n", distanceClosestIntersection)
}

func part2(panel *Panel) {
	shortestWireLengthIntersection := math.MaxInt

	for pointFirstWire, distFirstWire := range panel.first.coordinates {
		if distSecondWire, ok := panel.second.coordinates[pointFirstWire]; ok {
			distance := distSecondWire + distFirstWire
			shortestWireLengthIntersection = min(shortestWireLengthIntersection, distance)
		}
	}

	fmt.Printf("Part 1: %d\n", shortestWireLengthIntersection)
}

func parseInput(lines []string) (*Panel, error) {
	panel, err := NewPanel(lines)
	if err != nil {
		return nil, err
	}

	return panel, nil
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

	panel, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(panel)
	part2(panel)
}
