package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Assignament struct {
	start, end int
}

type AssignamentPair struct {
	first, second Assignament
}

func (a *AssignamentPair) Contains() bool {
	firstContainsSecond := a.first.start <= a.second.start && a.first.end >= a.second.end
	secondContainsFirst := a.second.start <= a.first.start && a.second.end >= a.first.end
	return firstContainsSecond || secondContainsFirst

}

func (a *AssignamentPair) Overlap() bool {
	firstOverlapsSecond := a.first.start <= a.second.start && a.second.start <= a.first.end
	secondOverlapsFirst := a.second.start <= a.first.start && a.first.start <= a.second.end
	return firstOverlapsSecond || secondOverlapsFirst

}

func NewAssignamentPair(line string) (*AssignamentPair, error) {
	var s1, e1, s2, e2 int
	_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &s1, &e1, &s2, &e2)
	if err != nil {
		return nil, err
	}

	return &AssignamentPair{
		first:  Assignament{start: s1, end: e1},
		second: Assignament{start: s2, end: e2},
	}, nil
}

func part1(assignaments []*AssignamentPair) {
	count := 0

	for _, assig := range assignaments {
		if assig.Contains() {
			count++
		}
	}

	fmt.Printf("Part 1: %d\n", count)
}

func part2(assignaments []*AssignamentPair) {
	count := 0

	for _, assig := range assignaments {
		if assig.Overlap() {
			count++
		}
	}

	fmt.Printf("Part 2: %d\n", count)

}

func ParseInput(lines []string) ([]*AssignamentPair, error) {
	assignaments := make([]*AssignamentPair, len(lines))
	for idx, line := range lines {
		currAssig, err := NewAssignamentPair(line)
		if err != nil {
			return nil, err
		}

		assignaments[idx] = currAssig
	}

	return assignaments, nil
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
	assignaments, err := ParseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(assignaments)
	part2(assignaments)
}
