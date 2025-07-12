package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Submarine struct {
	x, y int
	aim  int
}

func extractValue(command, prefix string) int {
	var val int
	fmt.Sscanf(command, prefix+" %d", &val)
	return val
}

func part1(commands []string) {
	s := Submarine{}

	for _, command := range commands {
		if strings.HasPrefix(command, "forward") {
			val := extractValue(command, "forward")
			s.x += val
		} else if strings.HasPrefix(command, "down") {
			val := extractValue(command, "down")
			s.y += val
		} else if strings.HasPrefix(command, "up") {
			val := extractValue(command, "up")
			s.y -= val
		}
	}

	fmt.Printf("Part 1: %d\n", s.x*s.y)

}

func part2(commands []string) {
	s := Submarine{}

	for _, command := range commands {
		if strings.HasPrefix(command, "forward") {
			val := extractValue(command, "forward")
			s.x += val
			s.y += s.aim * val
		} else if strings.HasPrefix(command, "down") {
			val := extractValue(command, "down")
			s.aim += val
		} else if strings.HasPrefix(command, "up") {
			val := extractValue(command, "up")
			s.aim -= val
		}
	}

	fmt.Printf("Part 2: %d\n", s.x*s.y)
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
	commands := strings.Split(string(fileContent), "\n")

	part1(commands)
	part2(commands)
}
