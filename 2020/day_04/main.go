package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	BIRT_YEAR       string = "byr"
	ISSUE_YEAR      string = "iyr"
	EXPIRATION_YEAR string = "eyr"
	HEIGHT          string = "hgt"
	HAIR_COLOR      string = "hcl"
	EYE_COLOR       string = "ecl"
	PASSPORT_ID     string = "pid"
	COUNTRY_ID      string = "cid"
)

var (
	color_regex    = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	validEyeColors = map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
)

func NewPassport(block string) (map[string]string, error) {
	lines := strings.Split(block, "\n")
	data := make(map[string]string)

	for _, line := range lines {
		parts := strings.Fields(strings.TrimSpace(line))
		for _, part := range parts {
			keyVal := strings.Split(part, ":")
			if len(keyVal) != 2 {
				return nil, fmt.Errorf("unexpected format")
			}
			data[keyVal[0]] = keyVal[1]
		}
	}

	return data, nil
}

func hasAllRequiredFields(pass map[string]string) bool {
	_, hasCountryID := pass[COUNTRY_ID]
	return len(pass) == 8 || (len(pass) == 7 && !hasCountryID)

}

func part1(passports []map[string]string) {
	validPassports := 0

	for _, pass := range passports {
		if hasAllRequiredFields(pass) {
			validPassports++
		}
	}

	fmt.Printf("Part 1: %d\n", validPassports)
}

func areYearsValid(pass map[string]string) bool {
	birthYear, err := strconv.Atoi(pass[BIRT_YEAR])
	if err != nil || birthYear < 1920 || birthYear > 2002 {
		return false
	}

	issueYear, err := strconv.Atoi(pass[ISSUE_YEAR])
	if err != nil || issueYear < 2010 || issueYear > 2020 {
		return false
	}

	expirationYear, err := strconv.Atoi(pass[EXPIRATION_YEAR])
	if err != nil || expirationYear < 2020 || expirationYear > 2030 {
		return false
	}

	return true
}

func isHeightValid(pass map[string]string) bool {
	height, ok := pass[HEIGHT]
	if !ok {
		return false
	}

	heightNum, err := strconv.Atoi(height[:len(height)-2])
	if err != nil {
		return false
	}

	if strings.HasSuffix(height, "cm") {
		return heightNum >= 150 && heightNum <= 193
	} else {
		return heightNum >= 59 && heightNum <= 76
	}
}

func areFieldsValid(pass map[string]string) bool {
	if !areYearsValid(pass) {
		return false
	}

	if !isHeightValid(pass) {
		return false
	}

	eyeColor, ok := pass[EYE_COLOR]
	if _, validColor := validEyeColors[eyeColor]; !ok || !validColor {
		return false
	}

	hairColor, ok := pass[HAIR_COLOR]
	if !ok || !color_regex.Match([]byte(hairColor)) {
		return false
	}

	passID, ok := pass[PASSPORT_ID]
	if !ok || len(passID) != 9 {
		return false
	}

	_, err := strconv.Atoi(passID)
	if err != nil {
		return false
	}

	return true
}

func part2(passports []map[string]string) {
	validPassports := 0

	for _, pass := range passports {
		if hasAllRequiredFields(pass) && areFieldsValid(pass) {
			validPassports++
		}
	}

	fmt.Printf("Part 2: %d\n", validPassports)
}

func parseInput(blocks []string) ([]map[string]string, error) {
	output := make([]map[string]string, len(blocks))

	for idx, block := range blocks {
		pass, err := NewPassport(block)
		if err != nil {
			return nil, err
		}

		output[idx] = pass
	}

	return output, nil
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
	input := strings.Split(string(fileContent), "\n\n")
	passports, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(passports)
	part2(passports)
}
