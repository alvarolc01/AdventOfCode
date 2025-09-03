package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	startNodePart1  = "AAA"
	targetNodePart1 = "ZZZ"
	nodeSuffixStart = "A"
	nodeSuffixEnd   = "Z"

	splitPartsCount = 2
)

type Directions struct {
	left, right string
}

func ParseLine(line string) (string, *Directions, error) {
	parts := strings.SplitN(line, "=", 2)
	if len(parts) != 2 {
		return "", nil, fmt.Errorf("invalid format")
	}

	key := strings.TrimSpace(parts[0])
	rest := strings.Trim(strings.TrimSpace(parts[1]), "()")
	dirs := strings.SplitN(rest, ",", 2)
	if len(dirs) != 2 {
		return "", nil, fmt.Errorf("invalid directions")
	}

	left := strings.TrimSpace(dirs[0])
	right := strings.TrimSpace(dirs[1])

	return key, &Directions{
		left:  left,
		right: right}, nil
}

func stepsToReach(nodes map[string]*Directions, steps, start string, complete func(string) bool) int {
	count, currStep := 0, 0
	currKey := start

	for !complete(currKey) {
		dirs := nodes[currKey]
		if steps[currStep] == 'L' {
			currKey = dirs.left
		} else {
			currKey = dirs.right
		}
		count++
		currStep = (currStep + 1) % len(steps)
	}

	return count
}

func part1(steps string, input map[string]*Directions) {
	requiredSteps := stepsToReach(input, steps, startNodePart1, func(key string) bool {
		return key == targetNodePart1
	})
	fmt.Printf("Part 1: %d\n", requiredSteps)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func part2(steps string, input map[string]*Directions) {
	currentNodes := []string{}
	for key := range input {
		if strings.HasSuffix(key, nodeSuffixStart) {
			currentNodes = append(currentNodes, key)
		}
	}

	output := stepsToReach(input, steps, currentNodes[0], func(node string) bool {
		return strings.HasSuffix(node, nodeSuffixEnd)
	})

	for _, start := range currentNodes[1:] {
		stepsFromStart := stepsToReach(input, steps, start, func(node string) bool {
			return strings.HasSuffix(node, nodeSuffixEnd)
		})
		output = lcm(output, stepsFromStart)
	}
	fmt.Printf("Part 2: %d\n", output)
}

func parseInput(input string) (string, map[string]*Directions, error) {
	parts := strings.SplitN(strings.TrimSpace(input), "\n\n", 2)
	if len(parts) != 2 {
		return "", nil, fmt.Errorf("unexpected format")
	}

	dirs := make(map[string]*Directions)
	lines := strings.Split(parts[1], "\n")
	for _, line := range lines {
		key, dir, err := ParseLine(line)
		if err != nil {
			return "", nil, err
		}
		dirs[key] = dir
	}

	return parts[0], dirs, nil

}

func main() {
	fileName := flag.String("file", "", "Path to input file")
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
	input := string(fileContent)

	steps, movements, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(steps, movements)
	part2(steps, movements)
}
