package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	lumberyardMarker rune = '#'
	groundMarker     rune = '.'
	treeMarker       rune = '|'

	stepsFirstPart  int = 10
	stepsSecondPart int = 1000000000
)

var directions = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func countSurrounding(currentMap []string, i, j int, marker rune) int {
	output := 0
	for _, d := range directions {
		ni, nj := i+d[0], j+d[1]
		if ni < 0 || ni >= len(currentMap) || nj < 0 || nj >= len(currentMap[i]) {
			continue
		}

		if rune(currentMap[ni][nj]) == marker {
			output++
		}
	}

	return output
}

func calculateNext(currentMap []string, i, j int) rune {
	cell := rune(currentMap[i][j])
	if cell == groundMarker {
		if countSurrounding(currentMap, i, j, treeMarker) >= 3 {
			return treeMarker
		}
		return groundMarker

	} else if cell == treeMarker {
		if countSurrounding(currentMap, i, j, lumberyardMarker) >= 3 {
			return lumberyardMarker
		}
		return treeMarker

	} else {
		lumberyards := countSurrounding(currentMap, i, j, lumberyardMarker)
		trees := countSurrounding(currentMap, i, j, treeMarker)
		if lumberyards >= 1 && trees >= 1 {
			return lumberyardMarker
		}
		return groundMarker
	}
}

func step(currentMap []string) []string {
	output := make([]string, len(currentMap))
	for i := range currentMap {
		row := make([]rune, len(currentMap[i]))
		for j := range currentMap[i] {
			row[j] = calculateNext(currentMap, i, j)
		}
		output[i] = string(row)
	}
	return output
}

func countResources(area []string) (trees, lumberyards int) {
	for _, row := range area {
		for _, c := range row {
			switch c {
			case treeMarker:
				trees++
			case lumberyardMarker:
				lumberyards++
			}
		}
	}
	return
}

func part1(currentMap []string) {
	for i := 0; i < stepsFirstPart; i++ {
		currentMap = step(currentMap)
	}

	trees, lumberyards := countResources(currentMap)
	fmt.Println("Part 1:", trees*lumberyards)

}

func part2(currentMap []string) {
	seen := make(map[string]int)
	mem := make(map[int]string)

	for i := 0; i < stepsSecondPart; i++ {
		key := strings.Join(currentMap, "\n")

		if prev, ok := seen[key]; ok {
			cycle := i - prev
			remaining := (stepsSecondPart - i) % cycle
			currentMap = strings.Split(mem[prev+remaining], "\n")
			break
		}
		seen[key] = i
		mem[i] = key
		currentMap = step(currentMap)
	}

	trees, lumberyards := countResources(currentMap)
	fmt.Println("Part 2:", trees*lumberyards)

}

func main() {
	fileName := flag.String("file", "", "Path to input file")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("Use --file to specify the input file path.")
		os.Exit(1)
	}

	data, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	part1(lines)
	part2(lines)
}
