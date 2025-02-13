package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func nextPassword(currentPassword string) string {
	idx := len(currentPassword) - 1
	nextPassword := []byte(currentPassword)

	for idx >= 0 {
		if nextPassword[idx] == 'z' {
			nextPassword[idx] = 'a'
			idx--
		} else {
			nextPassword[idx]++
			break
		}
	}

	return string(nextPassword)
}

func passwordHasIncreasingSubstring(password string) bool {
	for idx := 2; idx < len(password); idx++ {
		if password[idx]-password[idx-1] == 1 && password[idx-1]-password[idx-2] == 1 {
			return true
		}
	}
	return false
}

func passwordContainsInvalidChars(password string) bool {
	return strings.ContainsAny(password, "iol")
}

func passwordHasTwoDifferentPairs(password string) bool {
	countPairs := 0
	for idx := 1; idx < len(password); idx++ {
		if password[idx] == password[idx-1] {
			countPairs++
			idx++
		}

		if countPairs == 2 {
			return true
		}
	}
	return false
}

func isValidPassword(password string) bool {
	hasIncreasingSubstring := passwordHasIncreasingSubstring(password)
	containsInvalidChars := passwordContainsInvalidChars(password)
	hasTwoDifferentPairs := passwordHasTwoDifferentPairs(password)

	return hasIncreasingSubstring && !containsInvalidChars && hasTwoDifferentPairs
}

func getNextValidPassword(password string) string {
	for {
		password = nextPassword(password)
		if isValidPassword(password) {
			return password
		}
	}
}

func part1(currentPassword string) {
	nextValidPassword := getNextValidPassword(currentPassword)
	fmt.Printf("Part 1: %s\n", nextValidPassword)
}

func part2(currentPassword string) {
	nextValidPassword := getNextValidPassword(currentPassword)
	nextValidPassword = getNextValidPassword(nextValidPassword)
	fmt.Printf("Part 2: %s\n", nextValidPassword)
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
	input := string(fileContent)

	part1(input)
	part2(input)
}
