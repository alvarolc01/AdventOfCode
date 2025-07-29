package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ShinyGoldBag = "shiny gold"
)

func getBagContent(requirements map[string]map[string]int, targetBag string) int {
	requiredBags := 0

	for bagName, numBags := range requirements[targetBag] {
		requiredBags += numBags

		bagsInside := getBagContent(requirements, bagName)
		requiredBags += numBags * bagsInside
	}

	return requiredBags
}

func whichBagsContain(requirements map[string]map[string]int, targetBag string) map[string]bool {
	ans := make(map[string]bool)

	for key, val := range requirements {
		if _, ok := val[targetBag]; ok {
			ans[key] = true

			requireCurrent := whichBagsContain(requirements, key)
			for nameBag := range requireCurrent {
				ans[nameBag] = true
			}
		}
	}

	return ans
}

func part1(bagRequirements map[string]map[string]int) {
	bagsRequiringShinyGold := whichBagsContain(bagRequirements, ShinyGoldBag)
	differentColours := len(bagsRequiringShinyGold)

	fmt.Printf("Part 1: %d\n", differentColours)
}

func part2(bagRequirements map[string]map[string]int) {
	bagsInsideShinyGold := getBagContent(bagRequirements, ShinyGoldBag)

	fmt.Printf("Part 2: %d\n", bagsInsideShinyGold)
}

func GetBagRequirements(line string) (string, map[string]int, error) {
	fields := strings.Fields(line)
	baseBag := strings.Join(fields[:2], " ")

	requirements := make(map[string]int)
	if strings.Contains(line, "no other bags") {
		return baseBag, requirements, nil
	}

	for i := 3; i < len(fields); i++ {
		if strings.HasPrefix(fields[i], "bag") {
			numStr := fields[i-3]
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return "", nil, err
			}

			requiredBagName := strings.Join(fields[i-2:i], " ")
			requirements[requiredBagName] = num
		}
	}

	return baseBag, requirements, nil
}

func parseInput(lines []string) (map[string]map[string]int, error) {
	output := make(map[string]map[string]int)
	for _, line := range lines {
		badName, requirements, err := GetBagRequirements(line)
		if err != nil {
			return nil, err
		}
		output[badName] = requirements
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
	input := strings.Split(string(fileContent), "\n")

	bagRequirements, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(bagRequirements)
	part2(bagRequirements)
}
