package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Password struct {
	content             string
	firstVal, secondVal int
	ruleChar            rune
}

func NewPassword(line string) (*Password, error) {
	parts := strings.Split(line, ":")
	if len(parts) < 2 {
		return nil, fmt.Errorf("unexpected format")
	}

	content := strings.TrimSpace(parts[1])
	var firstVal, secondVal int
	var ruleChar rune

	_, err := fmt.Sscanf(parts[0], "%d-%d %c", &firstVal, &secondVal, &ruleChar)
	if err != nil {
		return nil, err
	}

	return &Password{
		content:   content,
		firstVal:  firstVal,
		secondVal: secondVal,
		ruleChar:  ruleChar,
	}, nil

}

func part1(passwords []*Password) {
	validPasswords := 0

	for _, pass := range passwords {
		countRepetitions := strings.Count(pass.content, string(pass.ruleChar))

		if countRepetitions >= pass.firstVal && countRepetitions <= pass.secondVal {
			validPasswords++
		}
	}

	fmt.Printf("Part 1: %d\n", validPasswords)
}

func part2(passwords []*Password) {
	validPasswords := 0

	for _, pass := range passwords {
		if (rune(pass.content[pass.firstVal-1]) == pass.ruleChar) != (rune(pass.content[pass.secondVal-1]) == pass.ruleChar) {
			validPasswords++
		}
	}

	fmt.Printf("Part 2: %d\n", validPasswords)
}

func parseInput(lines []string) ([]*Password, error) {
	output := make([]*Password, len(lines))
	for idx, line := range lines {
		pass, err := NewPassword(line)
		if err != nil {
			return nil, err
		}

		output[idx] = pass
	}

	return output, nil
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

	entries, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(entries)
	part2(entries)
}
