package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Replacement struct {
	original string
	replaced string
}

func part1(replacements []Replacement, originalString string) {
	possibleStrings := make(map[string]bool)
	for _, replacement := range replacements {

		re := regexp.MustCompile(replacement.original)
		matches := re.FindAllStringIndex(originalString, -1)
		for _, match := range matches {
			newString := originalString[:match[0]] + replacement.replaced + originalString[match[1]:]
			possibleStrings[newString] = true
		}
	}

	fmt.Printf("Part 1: %d\n", len(possibleStrings))
}

func part2(targetString string) {
	initialPolymers := 0
	for _, char := range targetString {
		if char >= 'A' && char <= 'Z' {
			initialPolymers++
		}
	}

	countY := strings.Count(targetString, "Y")
	countAr := strings.Count(targetString, "Ar")
	countRn := strings.Count(targetString, "Rn")

	totalSteps := initialPolymers - (countRn + countAr) - 2*countY - 1
	fmt.Printf("Part 2: %d\n", totalSteps)
}

func createReplacements(replacementsString string) []Replacement {
	var replacements []Replacement
	re := regexp.MustCompile(`(\w+) => (\w+)`)
	matches := re.FindAllStringSubmatch(replacementsString, -1)

	for _, match := range matches {
		newReplacement := Replacement{original: match[1], replaced: match[2]}
		replacements = append(replacements, newReplacement)
	}

	return replacements
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
	input := strings.Split(string(fileContent), "\n\n")

	replacements := createReplacements(input[0])

	part1(replacements, input[1])
	part2(input[1])
}
