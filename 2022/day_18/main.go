package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	x, y, z int
}

var directions = [6][3]int{
	{1, 0, 0},
	{-1, 0, 0},
	{0, 1, 0},
	{0, -1, 0},
	{0, 0, 1},
	{0, 0, -1},
}

func part1(cubes []*Point) {
	totalFaces := 0
	seen := make(map[Point]bool)
	for _, aCube := range cubes {
		currFaces := 6
		for _, dir := range directions {
			nextPos := Point{aCube.x + dir[0], aCube.y + dir[1], aCube.z + dir[2]}
			if seen[nextPos] {
				currFaces--
				totalFaces--
			}
		}
		seen[*aCube] = true
		totalFaces += currFaces
	}

	fmt.Printf("Part 1: %d\n", totalFaces)
}

func part2(cubes []*Point) {
	minX, minY, minZ := math.MaxInt, math.MaxInt, math.MaxInt
	maxX, maxY, maxZ := math.MinInt, math.MinInt, math.MinInt
	seen := make(map[Point]bool)
	for _, aCube := range cubes {
		minX = min(minX, aCube.x)
		minY = min(minY, aCube.y)
		minZ = min(minZ, aCube.z)
		maxX = max(maxX, aCube.x)
		maxY = max(maxY, aCube.y)
		maxZ = max(maxZ, aCube.z)
		seen[*aCube] = true
	}

	waterCubes := make(map[Point]bool)
	q := []Point{{minX, minY, minZ}}
	for len(q) > 0 {
		top := q[0]
		q = q[1:]

		if waterCubes[top] {
			continue
		}
		waterCubes[top] = true

		for _, dir := range directions {
			nextPos := Point{top.x + dir[0], top.y + dir[1], top.z + dir[2]}
			if nextPos.x < minX-1 || nextPos.x > maxX+1 ||
				nextPos.y < minY-1 || nextPos.y > maxY+1 ||
				nextPos.z < minZ-1 || nextPos.z > maxZ+1 {
				continue
			}

			if seen[nextPos] || waterCubes[nextPos] {
				continue
			}
			q = append(q, nextPos)
		}
	}

	totalFaces := 0
	for _, aCube := range cubes {
		for _, dir := range directions {
			pos := Point{aCube.x + dir[0], aCube.y + dir[1], aCube.z + dir[2]}
			if waterCubes[pos] {
				totalFaces++
			}
		}
	}
	fmt.Printf("Part 2: %d\n", totalFaces)
}

func parseInput(lines []string) ([]*Point, error) {
	output := make([]*Point, 0, len(lines))
	for _, aLine := range lines {
		currPoint := &Point{}
		_, err := fmt.Sscanf(aLine, "%d,%d,%d", &currPoint.x, &currPoint.y, &currPoint.z)
		if err != nil {
			return nil, err
		}
		output = append(output, currPoint)
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
	lines := strings.Split(string(fileContent), "\n")
	points, err := parseInput(lines)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}
	part1(points)
	part2(points)
}
