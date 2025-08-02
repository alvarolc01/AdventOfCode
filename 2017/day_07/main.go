package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Program struct {
	name     string
	size     int
	children []string
}

func NewProgram(line string) (*Program, error) {
	prog := Program{}

	parts := strings.Split(line, "->")

	_, err := fmt.Sscanf(strings.TrimSpace(parts[0]), "%s (%d)", &prog.name, &prog.size)
	if err != nil {
		return nil, err
	}

	if len(parts) > 1 {
		parts[1] = strings.Replace(parts[1], ",", " ", -1)
		fields := strings.Fields(parts[1])

		for _, field := range fields {
			prog.children = append(prog.children, field)
		}
	}

	return &prog, nil
}

func calcWeights(weights map[string]int, programs map[string]*Program, name string) int {
	if w, ok := weights[name]; ok {
		return w
	}

	prog := programs[name]
	base := prog.size

	for _, n := range prog.children {
		base += calcWeights(weights, programs, n)
	}

	weights[name] = base
	return base
}

func findImbalance(weights map[string]int, programs map[string]*Program, name string) int {
	prog := programs[name]

	childWeights := map[int][]string{}
	for _, child := range prog.children {
		imbalance := findImbalance(weights, programs, child)
		if imbalance != 0 {
			return imbalance
		}
		w := weights[child]
		childWeights[w] = append(childWeights[w], child)
	}

	if len(childWeights) <= 1 {
		return 0
	}

	var wrongWeight, correctWeight int
	for weight, names := range childWeights {
		if len(names) == 1 {
			wrongWeight = weight
		} else {
			correctWeight = weight
		}
	}

	diff := wrongWeight - correctWeight
	wrongChild := childWeights[wrongWeight][0]
	currentSize := programs[wrongChild].size

	return currentSize - diff
}

func part1(programs map[string]*Program) string {
	availablePrograms := make(map[string]bool)

	for key := range programs {
		availablePrograms[key] = true
	}

	for _, p := range programs {
		for _, ph := range p.children {
			delete(availablePrograms, ph)
		}
	}
	baseProgram := ""
	for key := range availablePrograms {
		baseProgram = key
	}

	fmt.Printf("Part 1: %s\n", baseProgram)
	return baseProgram
}

func part2(programs map[string]*Program, baseProgram string) {
	weights := make(map[string]int)
	calcWeights(weights, programs, baseProgram)
	imbalance := findImbalance(weights, programs, baseProgram)

	fmt.Printf("Part 2: %d\n", imbalance)

}

func parseInput(lines []string) (map[string]*Program, error) {
	output := make(map[string]*Program)
	for _, line := range lines {
		prog, err := NewProgram(line)
		if err != nil {
			return nil, err
		}
		output[prog.name] = prog
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

	progams, err := parseInput(input)
	if err != nil {
		fmt.Println("error parting input:", err)
		os.Exit(1)
	}

	baseProgram := part1(progams)
	part2(progams, baseProgram)
}
