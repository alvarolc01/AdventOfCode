package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var signalsToDigits = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

func part1(lines []string) {
	result := 0
	for _, line := range lines {
		numbers := strings.Split(line, "|")[1]
		fields := strings.Fields(numbers)
		for _, f := range fields {
			if len(f) == 2 || len(f) == 4 || len(f) == 3 || len(f) == 7 {
				result++
			}
		}
	}
	fmt.Printf("Part 1: %d\n", result)
}

func generateNumber(numbers, signals []string) int {
	segIndex := func(r rune) int {
		return int(r - 'a')
	}

	toBoolArray := func(pat string) [7]bool {
		var arr [7]bool
		for _, r := range pat {
			arr[segIndex(r)] = true
		}
		return arr
	}

	var four, seven [7]bool
	signalUsageCount := make(map[rune]int)
	for _, p := range numbers {
		for _, r := range p {
			signalUsageCount[r]++
		}
		switch len(p) {
		case 3:
			seven = toBoolArray(p)
		case 4:
			four = toBoolArray(p)
		}
	}

	correctWiring := make(map[rune]rune)

	for i := 0; i < 7; i++ {
		if seven[i] && !four[i] {
			correctWiring['a'] = rune('a' + i)
			break
		}
	}

	for r, n := range signalUsageCount {
		switch n {
		case 4:
			correctWiring['e'] = r
		case 6:
			correctWiring['b'] = r
		case 7:
			if !four[segIndex(r)] {
				correctWiring['g'] = r
			} else {
				correctWiring['d'] = r
			}
		case 8:
			if r != correctWiring['a'] {
				correctWiring['c'] = r
			}
		case 9:
			correctWiring['f'] = r
		}
	}

	initialSignalToNum := make(map[[7]bool]int)
	for requiredSignals, num := range signalsToDigits {
		var arr [7]bool
		for _, r := range requiredSignals {
			arr[segIndex(correctWiring[r])] = true
		}
		initialSignalToNum[arr] = num
	}

	var n int
	for _, inputSignal := range signals {
		n = 10*n + initialSignalToNum[toBoolArray(inputSignal)]
	}

	return n
}

func part2(lines []string) {
	result := 0
	for _, line := range lines {
		numbers := strings.Fields(strings.Split(line, "|")[0])
		signals := strings.Fields(strings.Split(line, "|")[1])

		result += generateNumber(numbers, signals)
	}
	fmt.Printf("Part 2: %d\n", result)
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

	part1(input)
	part2(input)
}
