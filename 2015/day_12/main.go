package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part1(input string) {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(input, -1)

	sumNumbers := 0
	for _, match := range matches {
		convertedNumber, err := strconv.Atoi(match)
		if err != nil {
			continue
		}
		sumNumbers += convertedNumber
	}

	fmt.Printf("Part 1: %d\n", sumNumbers)

}

func getSumValidNumbers(parsedInput interface{}) float64 {
	switch convertedInput := parsedInput.(type) {
	case float64:
		return convertedInput
	case []interface{}:
		sumNums := 0.0
		for _, val := range convertedInput {
			sumNums += getSumValidNumbers(val)
		}
		return sumNums
	case map[string]interface{}:
		sumNums := 0.0
		for _, val := range convertedInput {
			if val == "red" {
				return 0
			}
		}

		for _, val := range convertedInput {
			sumNums += getSumValidNumbers(val)
		}
		return sumNums
	default:
		return 0
	}
}

func part2(input string) {
	var parsedInput interface{}
	err := json.Unmarshal([]byte(input), &parsedInput)
	if err != nil {
		fmt.Printf("Error parsing input: %s\n", err)
		return
	}

	sumValidNumbers := getSumValidNumbers(parsedInput)

	fmt.Printf("Part 2: %d\n", int(sumValidNumbers))
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
