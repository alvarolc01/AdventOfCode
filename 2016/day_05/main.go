package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	PasswordLength       int    = 8
	IndexValueFirstPart  int    = 5
	IndexPosition        int    = 5
	IndexValueSecondPart int    = 6
	ExpectedPrefix       string = "00000"
)

func generateHash(doorId string, idx int) string {
	doorId += strconv.Itoa(idx)
	hash := md5.Sum([]byte(doorId))
	return hex.EncodeToString(hash[:])
}

func part1(doorID string) {
	password := ""

	for i := 0; len(password) != PasswordLength; i++ {
		hash := generateHash(doorID, i)
		if strings.HasPrefix(hash, ExpectedPrefix) {
			password += string(hash[IndexValueFirstPart])
		}
	}

	fmt.Printf("Part 1: %s\n", password)
}

func isValidHash(hash string, passwordCharacters map[int]rune) (bool, int) {
	hasPrefix := strings.HasPrefix(hash, ExpectedPrefix)
	if !hasPrefix {
		return false, 0
	}

	position := int(hash[IndexPosition]) - '0'
	if position >= PasswordLength || position < 0 {
		return false, 0
	}

	_, found := passwordCharacters[position]
	return !found, position
}

func getPasswordCharacters(doorID string) map[int]rune {
	passwordCharacters := make(map[int]rune)

	for i := 0; len(passwordCharacters) != PasswordLength; i++ {
		hash := generateHash(doorID, i)
		isValid, position := isValidHash(hash, passwordCharacters)
		if isValid {
			passwordCharacters[position] = rune(hash[IndexValueSecondPart])
		}
	}

	return passwordCharacters
}

func part2(doorID string) {
	passwordCharacters := getPasswordCharacters(doorID)

	password := ""
	for i := 0; i < PasswordLength; i++ {
		password += string(passwordCharacters[i])
	}

	fmt.Printf("Part 2: %s\n", password)
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
