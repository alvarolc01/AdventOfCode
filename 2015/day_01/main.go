package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func part1(input string) {
	floor := strings.Count(input, "(") - strings.Count(input, ")")

	fmt.Printf("Part1: %d\n", floor)
}

func part2(input string) {
	floor := 0
	var positionFirstTimeBasement int
	for idx, val := range input {
		if val == '(' {
			floor += 1
		} else if val == ')' {
			floor -= 1
		}

		if floor == -1 {
			positionFirstTimeBasement = idx + 1
			break
		}
	}
	fmt.Printf("Part2: %d\n", positionFirstTimeBasement)
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
		fmt.Printf("Error reading input file: %s\n", err)
		os.Exit(1)
	}
	input := string(fileContent)

	part1(input)
	part2(input)

}
