package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func countRepeatedCharacters(id string) (hasTwo, hasThree bool) {
	charCount := make([]int, 26)
	for _, c := range id {
		charCount[c-'a']++
	}

	for _, n := range charCount {
		if n == 2 {
			hasTwo = true
		} else if n == 3 {
			hasThree = true
		}
	}

	return
}

func part1(boxIDs []string) {
	countTwoRepetitions, countThreeRepetitions := 0, 0

	for _, id := range boxIDs {
		hasTwo, hasThree := countRepeatedCharacters(id)
		if hasTwo {
			countTwoRepetitions++
		}
		if hasThree {
			countThreeRepetitions++
		}
	}

	ans := countThreeRepetitions * countTwoRepetitions

	fmt.Printf("Part 1: %d\n", ans)

}

func commonLettersWithOneDiff(original, target string) (string, bool) {
	diffChars := 0
	var builder strings.Builder

	for i := range original {
		if target[i] != original[i] {
			diffChars++
			if diffChars > 1 {
				return "", false
			}
		} else {
			builder.WriteByte(original[i])
		}
	}

	return builder.String(), diffChars == 1
}

func part2(boxIDs []string) {
	var commonLetters string
	found := false

	for i := 0; i < len(boxIDs) && !found; i++ {
		for j := i + 1; j < len(boxIDs) && !found; j++ {
			if equalChars, oneCharDiff := commonLettersWithOneDiff(boxIDs[i], boxIDs[j]); oneCharDiff {
				found = true
				commonLetters = equalChars
			}
		}
	}

	fmt.Printf("Part 2: %s\n", commonLetters)
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
