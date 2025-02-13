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
	PrefixLengthFirstPart  int = 5
	PrefixLengthSecondPart int = 6
)

func GetMD5Hash(input string, num int) string {
	var sb strings.Builder
	sb.WriteString(input)
	sb.WriteString(strconv.Itoa(num))

	hash := md5.Sum([]byte(sb.String()))
	return hex.EncodeToString(hash[:])
}

func part1(input string) {
	currentNum := 0
	expectedPrefix := strings.Repeat("0", PrefixLengthFirstPart)
	currentHash := GetMD5Hash(input, currentNum)

	for !strings.HasPrefix(currentHash, expectedPrefix) {
		currentNum++
		currentHash = GetMD5Hash(input, currentNum)
	}

	fmt.Printf("Part 1: %d\n", currentNum)
}

func part2(input string) {
	currentNum := 0
	expectedPrefix := strings.Repeat("0", PrefixLengthSecondPart)
	currentHash := GetMD5Hash(input, currentNum)

	for !strings.HasPrefix(currentHash, expectedPrefix) {
		currentNum++
		currentHash = GetMD5Hash(input, currentNum)
	}

	fmt.Printf("Part 2: %d\n", currentNum)
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
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}
	input := string(fileContent)

	part1(input)
	part2(input)
}
