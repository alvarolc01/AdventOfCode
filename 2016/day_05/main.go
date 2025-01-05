package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
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
	position := int(hash[IndexPosition]) - '0'
	_, found := passwordCharacters[position]

	return !found && hasPrefix && position < PasswordLength, position
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

	inputStrings := readFile(fileName)

	part1(inputStrings)
	part2(inputStrings)

}
