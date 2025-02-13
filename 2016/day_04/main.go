package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

const (
	AlphabetSize     int    = 26
	PasswordLength   int    = 5
	SecondPartTarget string = "northpole"
)

func parseRoom(room string) (string, int, string) {
	re := regexp.MustCompile(`^(.*)-(\d+)\[(.+)\]$`)
	matches := re.FindStringSubmatch(room)

	if matches == nil {
		return "", 0, ""
	}

	encryptedName := matches[1]
	sectorId, err := strconv.Atoi(matches[2])
	if err != nil {
		return "", 0, ""
	}
	checksum := matches[3]

	return encryptedName, sectorId, checksum
}

func getFrequencyList(source string) map[rune]int {
	frequencyList := make(map[rune]int)

	for _, char := range source {
		if unicode.IsLetter(char) {
			frequencyList[char]++
		}
	}

	return frequencyList
}

func sortFrequencies(frequencyList map[rune]int) []rune {
	sortedCounts := make([]rune, 0, len(frequencyList))
	for letter := range frequencyList {
		sortedCounts = append(sortedCounts, letter)
	}

	sort.Slice(sortedCounts, func(i, j int) bool {
		if frequencyList[sortedCounts[i]] == frequencyList[sortedCounts[j]] {
			return sortedCounts[i] < sortedCounts[j]
		}
		return frequencyList[sortedCounts[i]] > frequencyList[sortedCounts[j]]
	})

	return sortedCounts
}

func getNMostFrequent(frequencyList map[rune]int) string {
	sortedList := sortFrequencies(frequencyList)
	password := ""

	for i := 0; i < PasswordLength && i < len(sortedList); i++ {
		password = password + string(sortedList[i])
	}

	return password
}

func isRoomReal(encryptedName, checksum string) bool {
	if encryptedName == "" {
		return false
	}
	frequencyList := getFrequencyList(encryptedName)
	obtainedPassword := getNMostFrequent(frequencyList)

	return obtainedPassword == checksum
}

func decryptChar(char rune, shiftForward int) rune {
	shiftedChar := (int(char) - 'a' + shiftForward) % AlphabetSize
	return rune(shiftedChar + 'a')
}

func decryptRoom(encryptedName string, shiftForward int) string {
	decryptedName := ""

	for _, char := range encryptedName {
		if char == '-' {
			decryptedName = decryptedName + " "
		} else {
			decryptedName = decryptedName + string(decryptChar(char, shiftForward))
		}
	}

	return decryptedName
}

func part1(input []string) {
	sumSectorsID := 0

	for _, room := range input {
		encryptedName, sectorID, checksum := parseRoom(room)
		if checksum != "" && isRoomReal(encryptedName, checksum) {
			sumSectorsID += sectorID
		}
	}

	fmt.Printf("Part 1: %d\n", sumSectorsID)
}

func part2(input []string) {
	foundSectorId := 0

	for roomNum := 0; roomNum < len(input) && foundSectorId == 0; roomNum++ {
		encryptedName, sectorID, _ := parseRoom(input[roomNum])
		decryptedName := decryptRoom(encryptedName, sectorID)

		if strings.Contains(strings.ToLower(decryptedName), SecondPartTarget) {
			foundSectorId = sectorID
		}
	}

	fmt.Printf("Part 2: %d\n", foundSectorId)
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
