package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Equation struct {
	ax, ay int
	bx, by int
	rx, ry int
}

func (e Equation) Solve() (int, int, bool) {
	det := e.ax*e.by - e.ay*e.bx
	if det == 0 {
		return 0, 0, false
	}

	xNum := e.rx*e.by - e.ry*e.bx
	yNum := e.ax*e.ry - e.ay*e.rx

	if xNum%det != 0 || yNum%det != 0 {
		return 0, 0, false
	}

	return xNum / det, yNum / det, true
}

func part1(equations []*Equation) {
	total := 0
	for _, eq := range equations {
		a, b, ok := eq.Solve()
		if ok {
			total += 3*a + b
		}
	}
	fmt.Printf("Part 1: %d\n", total)
}

func part2(equations []*Equation) {
	total := 0
	for _, eq := range equations {
		eq.rx += 10000000000000
		eq.ry += 10000000000000
		a, b, ok := eq.Solve()
		if ok {
			total += 3*a + b
		}
	}
	fmt.Printf("Part 2: %d\n", total)
}

func parseInput(blocks []string) ([]*Equation, error) {
	var result []*Equation
	for _, block := range blocks {
		lines := strings.Split(strings.TrimSpace(block), "\n")
		if len(lines) != 3 {
			return nil, fmt.Errorf("invalid block: %q", block)
		}
		var ax, ay, bx, by, rx, ry int
		if _, err := fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &ax, &ay); err != nil {
			return nil, err
		}
		if _, err := fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &bx, &by); err != nil {
			return nil, err
		}
		if _, err := fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &rx, &ry); err != nil {
			return nil, err
		}
		result = append(result, &Equation{ax, ay, bx, by, rx, ry})
	}
	return result, nil
}
func main() {
	fileName := flag.String("file", "", "Path to the file")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("file name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	content, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("error reading input file:", err)
		os.Exit(1)
	}

	blocks := strings.Split(string(content), "\n\n")
	equations, err := parseInput(blocks)
	if err != nil {
		fmt.Println("parsing error:", err)
		os.Exit(1)
	}

	part1(equations)
	part2(equations)

}
