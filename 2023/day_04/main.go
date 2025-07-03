package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	winningNumbers  []int
	selectedNumbers map[int]bool
}

func generateNumArray(str string) ([]int, error) {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str, -1)

	numbers := make([]int, len(matches))
	for i, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, err
		}
		numbers[i] = num
	}
	return numbers, nil
}

func NewCard(line string) (*Card, error) {
	parts := strings.Split(line, ":")
	cardParts := strings.Split(strings.TrimSpace(parts[1]), "|")

	winningNumbers, err := generateNumArray(cardParts[0])
	if err != nil {
		return nil, err
	}

	selectedNumbersList, err := generateNumArray(cardParts[1])
	if err != nil {
		return nil, err
	}

	selectedNumbers := make(map[int]bool)
	for _, n := range selectedNumbersList {
		selectedNumbers[n] = true
	}

	return &Card{
		winningNumbers:  winningNumbers,
		selectedNumbers: selectedNumbers,
	}, nil
}

func (c *Card) CountWinningNumbers() int {
	winningNumbers := 0

	for _, num := range c.winningNumbers {
		if c.selectedNumbers[num] {
			winningNumbers++
		}
	}

	return winningNumbers
}

func part1(cards []*Card) {
	sumCardScores := 0

	for _, card := range cards {
		winningNumbers := card.CountWinningNumbers()
		if winningNumbers > 0 {
			sumCardScores += int(math.Pow(float64(2), float64(winningNumbers-1)))
		}
	}

	fmt.Printf("Part 1: %d\n", sumCardScores)

}

func part2(cards []*Card) {
	countPlayedCards := 0
	timeWonCard := make([]int, len(cards))

	for idxCard, card := range cards {
		winningNumbers := card.CountWinningNumbers()
		timeWonCard[idxCard]++

		for idxFoudnCard := 0; idxFoudnCard < winningNumbers; idxFoudnCard++ {
			timeWonCard[idxCard+idxFoudnCard+1] += timeWonCard[idxCard]
		}

	}

	for _, n := range timeWonCard {
		countPlayedCards += n
	}

	fmt.Printf("Part 2: %d\n", countPlayedCards)

}

func parseInput(input []string) ([]*Card, error) {
	cards := make([]*Card, len(input))

	for idxLine, line := range input {
		currentCard, err := NewCard(line)
		if err != nil {
			return nil, err
		}

		cards[idxLine] = currentCard
	}

	return cards, nil
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

	cards, err := parseInput(input)
	if err != nil {
		fmt.Println("error reading parsing input:", err)
		os.Exit(1)
	}

	part1(cards)
	part2(cards)
}
