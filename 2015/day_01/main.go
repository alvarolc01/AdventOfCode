package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
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

func part1(input string) {
	floor := strings.Count(input, "(") - strings.Count(input, ")")

	fmt.Printf("Part1: %d\n", floor)
}
func part2(input string) {
	floor := 0
	for idx, val := range input {
		if val == '(' {
			floor += 1
		} else if val == ')' {
			floor -= 1
		}

		if floor == -1 {
			fmt.Printf("Part2: %d\n", idx+1)
			return
		}

	}
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
