package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

func isValidPassword(inputPassword string) bool {
	passwords := strings.Split(inputPassword, " ")
	foundWords := make(map[string]bool)

	for _, word := range passwords {
		if _, found := foundWords[word]; found {
			return false
		}
		foundWords[word] = true
	}

	return true
}

func part1(input []string) {
	validPasswords := 0

	for _, password := range input {
		if isValidPassword(password) {
			validPasswords++
		}
	}

	fmt.Printf("Part 1: %d \n", validPasswords)
}

func sortString(source string) string {
	runesWord := []rune(source)
	slices.Sort(runesWord)
	return string(runesWord)
}

func isValidPasswordAnagrams(inputPassword string) bool {
	passwords := strings.Split(inputPassword, " ")
	foundWords := make(map[string]bool)

	for _, word := range passwords {
		sortedWord := sortString(word)
		if _, found := foundWords[sortedWord]; found {
			return false
		}
		foundWords[sortedWord] = true
	}
	return true
}

func part2(input []string) {
	validPasswords := 0

	for _, password := range input {
		if isValidPasswordAnagrams(password) {
			validPasswords++
		}
	}

	fmt.Printf("Part 2: %d \n", validPasswords)
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
