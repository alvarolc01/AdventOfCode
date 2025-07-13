package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countOnesAtPos(lines []string, pos int) int {
	countOnes := 0

	for _, line := range lines {
		if line[pos] == '1' {
			countOnes++
		}
	}

	return countOnes
}

func binaryStringToInt(s string) int64 {
	num, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return 0
	}
	return num
}

func getRate(measurements []string, comparator func(ones, zeroes int) byte) int64 {
	rate := ""
	for idx := range len(measurements[0]) {
		onesAtPos := countOnesAtPos(measurements, idx)
		zeroesAtPos := len(measurements) - onesAtPos
		rate += string(comparator(onesAtPos, zeroesAtPos))
	}
	return binaryStringToInt(rate)
}

func part1(measurements []string) {
	gamma := getRate(measurements, func(ones, zeroes int) byte {
		if ones >= zeroes {
			return '1'
		}
		return '0'
	})

	epsilon := getRate(measurements, func(ones, zeroes int) byte {
		if ones < zeroes {
			return '1'
		}
		return '0'
	})

	fmt.Printf("Part 1: %d\n", gamma*epsilon)
}

func filterList(measurements []string, pos int, val byte) []string {
	var output []string

	for _, line := range measurements {
		if line[pos] == val {
			output = append(output, line)
		}
	}

	return output
}

func getLifeSupportRating(measurements []string, comparator func(ones, zeroes int) byte) int64 {
	validMeasurements := measurements
	for idx := 0; idx < len(validMeasurements[0]) && len(validMeasurements) > 1; idx++ {
		ones := countOnesAtPos(validMeasurements, idx)
		zeroes := len(validMeasurements) - ones
		validMeasurements = filterList(validMeasurements, idx, comparator(ones, zeroes))
	}

	return binaryStringToInt(validMeasurements[0])
}

func part2(measurements []string) {
	oxygen := getLifeSupportRating(measurements, func(ones, zeroes int) byte {
		if ones >= zeroes {
			return '1'
		}
		return '0'
	})

	co2 := getLifeSupportRating(measurements, func(ones, zeroes int) byte {
		if ones < zeroes {
			return '1'
		}
		return '0'
	})

	fmt.Printf("Part 2: %d\n", co2*oxygen)
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
