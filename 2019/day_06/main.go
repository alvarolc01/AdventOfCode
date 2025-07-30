package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	CentreOfMass          = "COM"
	OriginOrbitalTransfer = "YOU"
	EndOrbitalTransfer    = "SAN"
)

func part1(orbits map[string][]string) {
	totalOrbitsSum := 0
	currDepth := 0
	queue := []string{CentreOfMass}

	for len(queue) > 0 {
		nextDistance := []string{}

		for _, planet := range queue {
			totalOrbitsSum += currDepth
			nextDistance = append(nextDistance, orbits[planet]...)
		}

		currDepth++
		queue = nextDistance
	}

	fmt.Printf("Part 1: %d\n", totalOrbitsSum)
}

func getPathTo(orbits map[string]string, planet string) []string {
	output := []string{}

	for currPlanet := planet; currPlanet != CentreOfMass; currPlanet = orbits[currPlanet] {
		output = append(output, currPlanet)
	}
	output = append(output, CentreOfMass)

	slices.Reverse(output)
	return output
}

func part2(orbits map[string]string) {
	totalOrbitsSum := 0

	pathToOri := getPathTo(orbits, OriginOrbitalTransfer)
	pathToEnd := getPathTo(orbits, EndOrbitalTransfer)

	matchingOrbits := 0
	for matchingOrbits+1 < len(pathToEnd) && matchingOrbits+1 < len(pathToOri) && pathToEnd[matchingOrbits+1] == pathToOri[matchingOrbits+1] {
		matchingOrbits++
	}

	totalOrbitsSum = len(pathToEnd) + len(pathToOri) - 2*(2+matchingOrbits)
	fmt.Printf("Part 2: %d\n", totalOrbitsSum)
}

func parseInput(lines []string) (map[string][]string, map[string]string, error) {
	fromCentre := make(map[string][]string)
	toCentre := make(map[string]string)

	for _, line := range lines {
		parts := strings.Split(line, ")")
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("unexpected format")
		}

		parent, child := parts[0], parts[1]
		if _, ok := fromCentre[parent]; !ok {
			fromCentre[parent] = make([]string, 0)
		}
		fromCentre[parent] = append(fromCentre[parent], child)
		toCentre[child] = parent
	}

	return fromCentre, toCentre, nil
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

	fromCentre, toCentre, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(fromCentre)
	part2(toCentre)
}
