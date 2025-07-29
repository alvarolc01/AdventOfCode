package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type AnswerGroup struct {
	answerCount map[rune]int
	numPeople   int
}

func NewGroupAnswer(block string) (*AnswerGroup, error) {
	output := make(map[rune]int)
	lines := strings.Split(block, "\n")

	for _, line := range lines {
		for _, char := range line {
			output[char]++
		}
	}

	return &AnswerGroup{
		answerCount: output,
		numPeople:   len(lines),
	}, nil
}

func part1(answerGroups []*AnswerGroup) {
	sumVotedMeasures := 0

	for _, group := range answerGroups {
		sumVotedMeasures += len(group.answerCount)
	}

	fmt.Printf("Part 1: %d\n", sumVotedMeasures)
}

func part2(answerGroups []*AnswerGroup) {
	countUnanimousVotes := 0

	for _, group := range answerGroups {
		for _, peopleInFavour := range group.answerCount {
			if peopleInFavour == group.numPeople {
				countUnanimousVotes++
			}
		}
	}

	fmt.Printf("Part 2: %d\n", countUnanimousVotes)
}

func parseInput(blocks []string) ([]*AnswerGroup, error) {
	output := make([]*AnswerGroup, len(blocks))
	for idx, block := range blocks {
		groupAnswer, err := NewGroupAnswer(block)
		if err != nil {
			return nil, err
		}

		output[idx] = groupAnswer
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
	input := strings.Split(string(fileContent), "\n\n")

	entries, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(entries)
	part2(entries)
}
