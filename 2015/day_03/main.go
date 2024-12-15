package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func readFile(fileName *string) string {
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read file content: %s", err)
		os.Exit(1)
	}

	return string(content)

}

type Position struct {
	positionXAxis, positionYAxis int
}

func nextPosition(movement rune, currentPosition Position) Position {
	if movement == '^' {
		currentPosition.positionYAxis++
	} else if movement == '<' {
		currentPosition.positionXAxis--
	} else if movement == '>' {
		currentPosition.positionXAxis++
	} else if movement == 'v' {
		currentPosition.positionYAxis--
	}

	return currentPosition
}

func part1(input string) {
	visited := map[Position]bool{Position{positionXAxis: 0, positionYAxis: 0}: true}

	santaPosition := Position{positionXAxis: 0, positionYAxis: 0}

	for _, movement := range input {
		santaPosition = nextPosition(movement, santaPosition)
		visited[santaPosition] = true
	}

	fmt.Printf("Part1: %d\n", len(visited))
}

func part2(input string) {
	visited := map[Position]bool{Position{positionXAxis: 0, positionYAxis: 0}: true}

	santaLocation := Position{positionXAxis: 0, positionYAxis: 0}
	robotLocation := Position{positionXAxis: 0, positionYAxis: 0}

	for idx, movement := range input {
		if idx%2 == 0 {
			santaLocation = nextPosition(movement, santaLocation)
			visited[santaLocation] = true
		} else {
			robotLocation = nextPosition(movement, robotLocation)
			visited[robotLocation] = true
		}
	}

	fmt.Printf("Part2: %d\n", len(visited))
}
func main() {
	fileName := flag.String("file", "", "Path to the file to read")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("File name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	inputString := readFile(fileName)

	part1(inputString)
	part2(inputString)

}
