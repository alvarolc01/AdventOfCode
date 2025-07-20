package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Toboggan struct {
	AreaMap []string
}

func (t *Toboggan) TreesInRoute(dx, dy int) int {
	trees := 0
	x, y := 0, 0

	for x < len(t.AreaMap) {
		if t.AreaMap[x][y] == '#' {
			trees++
		}
		x += dx
		y = (y + dy) % len(t.AreaMap[0])
	}

	return trees
}

func part1(toboggan *Toboggan) {
	trees := toboggan.TreesInRoute(1, 3)
	fmt.Printf("Part 1: %d\n", trees)
}

func part2(toboggan *Toboggan) {
	treesMult := 1
	steps := [][2]int{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}

	for _, step := range steps {
		treesMult *= toboggan.TreesInRoute(step[0], step[1])
	}

	fmt.Printf("Part 2: %d\n", treesMult)
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

	toboggan := &Toboggan{
		AreaMap: input,
	}

	part1(toboggan)
	part2(toboggan)
}
