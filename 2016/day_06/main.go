package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func getCharCountPerPosition(lines []string) [][]int {
	charCountPerPosition := make([][]int, len(lines[0]))
	for i := range charCountPerPosition {
		charCountPerPosition[i] = make([]int, 26)
	}

	for _, line := range lines {
		for idxChar, char := range line {
			charCountPerPosition[idxChar][int(char-'a')]++
		}
	}

	return charCountPerPosition
}

func getMostFrequentChar(frequencyList []int) rune {
	maxFrequencyPos := 0
	for idx, count := range frequencyList {
		if count > frequencyList[maxFrequencyPos] {
			maxFrequencyPos = idx
		}
	}

	return rune(maxFrequencyPos + 'a')
}

func getLeastFrequentChar(frequencyList []int) rune {
	minFrequencyPos := 0
	foundFirstChar := false
	for idx, count := range frequencyList {
		fewerApparitions := count != 0 && count < frequencyList[minFrequencyPos]
		if !foundFirstChar || fewerApparitions {
			minFrequencyPos = idx
			foundFirstChar = true
		}
	}

	return rune(minFrequencyPos + 'a')
}

func part1(input []string) {
	charCountPerPos := getCharCountPerPosition(input)

	mostFrequentCharsWord := ""
	for _, charsCount := range charCountPerPos {
		mostRepeatedChar := getMostFrequentChar(charsCount)
		mostFrequentCharsWord += string(mostRepeatedChar)
	}

	fmt.Printf("Part 1: %s\n", mostFrequentCharsWord)
}

func part2(input []string) {
	chatCountPerPos := getCharCountPerPosition(input)

	leastFrequentCharsWord := ""
	for _, charsCount := range chatCountPerPos {
		lastRepeatedChar := getLeastFrequentChar(charsCount)
		leastFrequentCharsWord += string(lastRepeatedChar)

	}
	fmt.Printf("Part 2: %s\n", leastFrequentCharsWord)
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

	part1(input)
	part2(input)

}
