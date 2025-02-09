package main

import (
	"flag"
	"fmt"
	"os"
)

type Position struct {
	positionXAxis, positionYAxis int
}

func (p *Position) Move(movement rune) {
	if movement == '^' {
		p.positionYAxis++
	} else if movement == '<' {
		p.positionXAxis--
	} else if movement == '>' {
		p.positionXAxis++
	} else if movement == 'v' {
		p.positionYAxis--
	}
}

func part1(input string) {
	visited := map[Position]bool{Position{positionXAxis: 0, positionYAxis: 0}: true}

	santaPosition := Position{positionXAxis: 0, positionYAxis: 0}

	for _, movement := range input {
		santaPosition.Move(movement)
		visited[santaPosition] = true
	}

	fmt.Printf("Part 1: %d\n", len(visited))
}

func part2(input string) {
	visited := map[Position]bool{Position{positionXAxis: 0, positionYAxis: 0}: true}

	santaLocation := Position{positionXAxis: 0, positionYAxis: 0}
	robotLocation := Position{positionXAxis: 0, positionYAxis: 0}

	for idx, movement := range input {
		if idx%2 == 0 {
			santaLocation.Move(movement)
			visited[santaLocation] = true
		} else {
			robotLocation.Move(movement)
			visited[robotLocation] = true
		}
	}

	fmt.Printf("Part 2: %d\n", len(visited))
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
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}
	input := string(fileContent)

	part1(input)
	part2(input)

}
