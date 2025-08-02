package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

const AllowedSumDistances = 10000

type Pos struct {
	x, y int
}

func (p Pos) ManhattanDistance(i, j int) int {
	return int(math.Abs(float64(p.x-i)) + math.Abs(float64(p.y-j)))
}

func getEdges(points []Pos) (int, int, int, int) {
	maxX, minX, maxY, minY := math.MinInt, math.MaxInt, math.MinInt, math.MaxInt

	for _, point := range points {
		maxX = max(maxX, point.x)
		minX = min(minX, point.x)
		maxY = max(maxY, point.y)
		minY = min(minY, point.y)
	}

	return maxX, minX, maxY, minY
}

func part1(points []Pos) {
	count := make(map[int]int)
	for i := range points {
		count[i] = 0
	}

	maxX, minX, maxY, minY := getEdges(points)

	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			closestIdx := -1
			minDist := math.MaxInt

			tie := false
			for idx, point := range points {
				dist := point.ManhattanDistance(i, j)
				if dist < minDist {
					minDist = dist
					closestIdx = idx
					tie = false
				} else if dist == minDist {
					tie = true
				}
			}

			if tie {
				continue
			}

			if i == maxX || i == minX || i == maxY || i == minY {
				delete(count, closestIdx)
			} else if _, ok := count[closestIdx]; ok {
				count[closestIdx]++
			}
		}
	}

	maxArea := 0
	for _, val := range count {
		if val > maxArea {
			maxArea = val
		}
	}

	fmt.Printf("Part 1: %d\n", maxArea)
}

func part2(points []Pos) {
	count := 0
	maxX, minX, maxY, minY := getEdges(points)

	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			sumDistances := 0
			for _, point := range points {
				dist := point.ManhattanDistance(i, j)
				sumDistances += dist
			}

			if sumDistances < AllowedSumDistances {
				count++
			}
		}
	}

	fmt.Printf("Part 2: %d\n", count)
}

func parseInput(lines []string) ([]Pos, error) {
	output := make([]Pos, len(lines))
	for i, line := range lines {
		var x, y int
		_, err := fmt.Sscanf(line, "%d, %d", &x, &y)
		if err != nil {
			return nil, err
		}
		output[i] = Pos{x, y}
	}
	return output, nil
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
	points, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(points)
	part2(points)
}
