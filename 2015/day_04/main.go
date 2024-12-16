package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
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

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func StarsWithNZeros(attempt string, numZeros int) bool {
	return attempt[0:numZeros] == strings.Repeat("0", numZeros)
}

func part1(input string) {
	addedNum := 0
	currentAttempt := fmt.Sprintf("%s%d", input, addedNum)
	for !StarsWithNZeros(GetMD5Hash(currentAttempt), 5) {
		addedNum++
		currentAttempt = fmt.Sprintf("%s%d", input, addedNum)
	}

	fmt.Printf("Part 1: %d\n", addedNum)
}

func part2(input string) {
	addedNum := 0
	currentAttempt := fmt.Sprintf("%s%d", input, addedNum)
	for !StarsWithNZeros(GetMD5Hash(currentAttempt), 6) {
		addedNum++
		currentAttempt = fmt.Sprintf("%s%d", input, addedNum)
	}

	fmt.Printf("Part2: %d\n", addedNum)
}

func main() {
	fileName := flag.String("file", "", "Path to the file to read")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("File name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	inputString := readFile(fileName)

	part1(inputString)
	part2(inputString)

}
