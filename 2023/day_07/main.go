package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	FiveOfAKind = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

var value = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type Hand struct {
	cards string
	bid   int
}

func (h *Hand) Type(useJokers bool) int {
	labels := make(map[rune]int)
	for _, card := range h.cards {
		labels[card]++
	}

	jokers := 0
	if useJokers {
		jokers = labels['J']
		delete(labels, 'J')
	}

	repetitions := []int{}
	for _, v := range labels {
		repetitions = append(repetitions, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(repetitions)))

	if len(repetitions) == 0 {
		return FiveOfAKind
	}

	if useJokers {
		repetitions[0] += jokers
	}

	switch {
	case repetitions[0] == 5:
		return FiveOfAKind
	case repetitions[0] == 4:
		return FourOfAKind
	case repetitions[0] == 3 && len(repetitions) > 1 && repetitions[1] == 2:
		return FullHouse
	case repetitions[0] == 3:
		return ThreeOfAKind
	case repetitions[0] == 2 && len(repetitions) > 1 && repetitions[1] == 2:
		return TwoPair
	case repetitions[0] == 2:
		return OnePair
	default:
		return HighCard
	}
}

func NewHand(line string) (*Hand, error) {
	parts := strings.Fields(line)
	if len(parts) != 2 {
		return nil, fmt.Errorf("expected two parts per line")
	}

	cards := parts[0]
	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	return &Hand{
		cards: cards,
		bid:   bid,
	}, nil
}

func sortHands(hands []*Hand, useJokers bool, getValueFunc func(byte) int) {
	sort.Slice(hands, func(i, j int) bool {
		typeI := hands[i].Type(useJokers)
		typeJ := hands[j].Type(useJokers)
		if typeI < typeJ {
			return true
		} else if typeJ < typeI {
			return false
		}

		for pos := range hands[i].cards {
			valCardI := getValueFunc(hands[i].cards[pos])
			valCardJ := getValueFunc(hands[j].cards[pos])
			if valCardI > valCardJ {
				return true
			} else if valCardJ > valCardI {
				return false
			}
		}

		return true
	})
}

func part1(hands []*Hand) {
	sortHands(hands, false, func(card byte) int {
		return value[card]
	})

	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += hand.bid * (len(hands) - i)
	}

	fmt.Printf("Part 1: %d\n", totalWinnings)
}

func part2(hands []*Hand) {
	sortHands(hands, true, func(card byte) int {
		if card == 'J' {
			return 1
		}
		return value[card]
	})

	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += hand.bid * (len(hands) - i)
	}

	fmt.Printf("Part 2: %d\n", totalWinnings)
}

func parseInput(input []string) ([]*Hand, error) {
	output := make([]*Hand, len(input))

	for idx, line := range input {
		hand, err := NewHand(line)
		if err != nil {
			return nil, err
		}

		output[idx] = hand
	}

	return output, nil
}

func main() {
	fileName := flag.String("file", "", "Path to input file")
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

	hands, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(hands)
	part2(hands)
}
