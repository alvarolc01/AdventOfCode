package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

type Node struct {
	x, y   int
	vx, vy int
}

func parseInput(lines []string) ([]*Node, error) {
	nodes := make([]*Node, 0, len(lines))
	for _, line := range lines {

		var x, y, vx, vy int
		_, err := fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &x, &y, &vx, &vy)
		if err != nil {
			return nil, fmt.Errorf("invalid line: %q", line)
		}

		nodes = append(nodes, &Node{x: x, y: y, vx: vx, vy: vy})
	}
	return nodes, nil
}

func step(nodes []*Node, forward bool) {
	for _, n := range nodes {
		if forward {
			n.x += n.vx
			n.y += n.vy
		} else {
			n.x -= n.vx
			n.y -= n.vy
		}
	}
}

func getBoundaries(nodes []*Node) (int, int, int, int) {
	left, top := math.MaxInt, math.MaxInt
	right, bottom := math.MinInt, math.MinInt
	for _, n := range nodes {
		if n.x < left {
			left = n.x
		}
		if n.x > right {
			right = n.x
		}
		if n.y < top {
			top = n.y
		}
		if n.y > bottom {
			bottom = n.y
		}
	}
	return left, top, right, bottom
}

func area(nodes []*Node) int {
	l, t, r, b := getBoundaries(nodes)
	return (r - l) * (b - t)
}

func printGraph(nodes []*Node) {
	left, top, right, bottom := getBoundaries(nodes)
	found := map[[2]int]bool{}
	for _, n := range nodes {
		found[[2]int{n.x, n.y}] = true
	}
	fmt.Println("Part 1:")
	for y := top; y <= bottom; y++ {
		for x := left; x <= right; x++ {
			if found[[2]int{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func part1(nodes []*Node) {
	prevArea := math.MaxInt
	seconds := 0

	for {
		step(nodes, true)
		seconds++
		a := area(nodes)
		if a > prevArea {
			step(nodes, false)
			seconds--
			printGraph(nodes)
			fmt.Printf("Part 2: %d\n", seconds)
			return
		}
		prevArea = a
	}
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
	nodes, err := parseInput(lines)
	if err != nil {
		fmt.Println("Parse error:", err)
		os.Exit(1)
	}

	part1(nodes)
}
