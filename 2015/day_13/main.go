package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type PersonChanges struct {
	values map[string]int
}

func NewPersonChanges() *PersonChanges {
	return &PersonChanges{
		values: make(map[string]int),
	}
}

func (pc *PersonChanges) addChange(destination string, value int) {
	pc.values[destination] = value
}

func (pc *PersonChanges) getHappiness(destination string) int {
	return pc.values[destination]
}

type HappinessChanges struct {
	personalChanges map[string]*PersonChanges
}

func NewHappinessChanges() *HappinessChanges {
	return &HappinessChanges{
		personalChanges: make(map[string]*PersonChanges),
	}
}

func (h *HappinessChanges) addPersonalChange(source, destination string, change int) {
	if _, exists := h.personalChanges[source]; !exists {
		h.personalChanges[source] = NewPersonChanges()
	}

	h.personalChanges[source].addChange(destination, change)
}

func (h *HappinessChanges) calculateHappinessChange(seatingArrangement []string) int {
	totalChanges := 0
	num := len(seatingArrangement)

	for idx := 0; idx < num; idx++ {
		idxLeft := (idx - 1 + num) % num
		idxRight := (idx + 1) % num
		currentGuest := seatingArrangement[idx]

		totalChanges += h.personalChanges[currentGuest].getHappiness(seatingArrangement[idxLeft])
		totalChanges += h.personalChanges[currentGuest].getHappiness(seatingArrangement[idxRight])

	}

	return totalChanges
}

func backtrackingHappiness(hc *HappinessChanges, seatingArrangement []string, seatedPeople map[string]bool) int {
	if len(hc.personalChanges) == len(seatingArrangement) {
		return hc.calculateHappinessChange(seatingArrangement)
	}

	happiness := 0
	for key := range hc.personalChanges {
		if !seatedPeople[key] {
			seatedPeople[key] = true
			maxHappinessCurrentSeating := backtrackingHappiness(hc, append(seatingArrangement, key), seatedPeople)
			happiness = max(happiness, maxHappinessCurrentSeating)
			delete(seatedPeople, key)
		}
	}
	return happiness
}

func part1(input *HappinessChanges) {
	maxHappiness := backtrackingHappiness(input, []string{}, map[string]bool{})
	fmt.Printf("Part 1: %d\n", maxHappiness)
}

func part2(input *HappinessChanges) {
	for person := range input.personalChanges {
		input.addPersonalChange("myself", person, 0)
		input.addPersonalChange(person, "myself", 0)
	}
	maxHappiness := backtrackingHappiness(input, []string{}, map[string]bool{})
	fmt.Printf("Part 2: %d\n", maxHappiness)
}

func parseInput(input []string) *HappinessChanges {
	hc := NewHappinessChanges()
	re := regexp.MustCompile(`(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+).`)

	for _, line := range input {
		match := re.FindStringSubmatch(line)
		changeVal, err := strconv.Atoi(match[3])
		if err != nil {
			continue
		}
		if match[2] == "lose" {
			changeVal = -changeVal
		}
		hc.addPersonalChange(match[1], match[4], changeVal)
	}

	return hc
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

	happinessChanges := parseInput(input)
	part1(happinessChanges)
	part2(happinessChanges)
}
