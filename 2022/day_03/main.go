package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func commonChars(rucksacks []string) []byte {
	count := make(map[byte]int)
	for i := range rucksacks[0] {
		count[rucksacks[0][i]] = 1
	}

	for i := 1; i < len(rucksacks); i++ {
		for _, char := range rucksacks[i] {
			val := count[byte(char)]
			if val == i {
				count[byte(char)] = i + 1
			} else if val < i {
				delete(count, byte(char))
			}
		}
	}

	var output []byte
	for key, val := range count {
		if val == len(rucksacks) {
			output = append(output, key)
		}
	}
	return output
}

func getScoreChar(char byte) int {
	if char >= 'a' && char <= 'z' {
		return int(char-'a') + 1
	} else if char >= 'A' && char <= 'Z' {
		return int(char-'A') + 27
	}
	return 0
}

func part1(rucksacks []string) {
	prioritiesSum := 0

	for _, rucksack := range rucksacks {
		parts := []string{rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]}
		comChars := commonChars(parts)
		if len(comChars) != 1 {
			continue
		}

		prioritiesSum += getScoreChar(comChars[0])
	}

	fmt.Printf("Part 1: %d\n", prioritiesSum)
}

func part2(rucksacks []string) {
	prioritiesSum := 0

	for idx := 0; idx < len(rucksacks); idx += 3 {
		comChars := commonChars(rucksacks[idx : idx+3])
		if len(comChars) != 1 {
			continue
		}

		prioritiesSum += getScoreChar(comChars[0])
	}

	fmt.Printf("Part 2: %d\n", prioritiesSum)
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

	part1(input)
	part2(input)
}
