package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Graph struct {
	adjacencyList map[string][]Edge
}

type Edge struct {
	destination string
	distance    int
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[string][]Edge),
	}
}

func (g *Graph) AddEdge(source, destination string, distance int) {
	if _, ok := g.adjacencyList[source]; !ok {
		g.adjacencyList[source] = make([]Edge, 0)
	}
	g.adjacencyList[source] = append(g.adjacencyList[source], Edge{destination, distance})

	if _, ok := g.adjacencyList[destination]; !ok {
		g.adjacencyList[destination] = make([]Edge, 0)
	}
	g.adjacencyList[destination] = append(g.adjacencyList[destination], Edge{source, distance})
}

func (g *Graph) backtrackingRoutes(currentLocation string, visited map[string]bool, currentDistance int, updateResponse func(int, int) int) int {
	if len(g.adjacencyList) == len(visited) {
		return currentDistance
	}

	traveledSoFar := 0
	reachableLocations, _ := g.adjacencyList[currentLocation]
	for _, travel := range reachableLocations {
		if _, ok := visited[travel.destination]; !ok {
			visited[travel.destination] = true
			addingLocationDistance := g.backtrackingRoutes(travel.destination, visited, currentDistance+travel.distance, updateResponse)
			delete(visited, travel.destination)
			traveledSoFar = updateResponse(traveledSoFar, addingLocationDistance)
		}
	}

	return traveledSoFar
}

func part1(graph *Graph) {
	getMinDistance := func(currentMin, routeDistance int) int {
		if currentMin == 0 {
			return routeDistance
		}
		return min(currentMin, routeDistance)
	}

	minDistanceRoute := 0
	for startingLocation := range graph.adjacencyList {
		visited := map[string]bool{startingLocation: true}
		minDistanceFromLocation := graph.backtrackingRoutes(startingLocation, visited, 0, getMinDistance)
		minDistanceRoute = getMinDistance(minDistanceRoute, minDistanceFromLocation)
	}

	fmt.Printf("Part 1: %d\n", minDistanceRoute)
}

func part2(graph *Graph) {
	getMaxDistance := func(currentMax, routeDistance int) int {
		if currentMax == 0 {
			return routeDistance
		}
		return max(currentMax, routeDistance)
	}

	maxDistanceRoute := 0
	for startingLocation := range graph.adjacencyList {
		visited := map[string]bool{startingLocation: true}
		maxDistanceFromLocation := graph.backtrackingRoutes(startingLocation, visited, 0, getMaxDistance)
		maxDistanceRoute = getMaxDistance(maxDistanceRoute, maxDistanceFromLocation)
	}

	fmt.Printf("Part 2: %d\n", maxDistanceRoute)
}

func parseInput(input []string) *Graph {
	result := NewGraph()
	re := regexp.MustCompile(`^(\w+) to (\w+) = (\d+)$`)
	for _, line := range input {
		matches := re.FindStringSubmatch(line)
		distance, err := strconv.Atoi(matches[3])
		if err != nil {
			continue
		}
		result.AddEdge(matches[1], matches[2], distance)
	}
	return result
}

func main() {
	fileName := flag.String("file", "", "Path to the file to read")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("File name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	fileContent, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}
	input := strings.Split(string(fileContent), "\n")

	validTravels := parseInput(input)

	part1(validTravels)
	part2(validTravels)
}
