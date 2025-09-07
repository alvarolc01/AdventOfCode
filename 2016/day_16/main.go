package main

import (
	"fmt"
	"slices"
)

const (
	diskSizePartOne int = 272
	diskSizePartTwo int = 35651584
)

func generateDragonCurve(initialDiskState string) string {
	copyOriginal := make([]rune, 0, len(initialDiskState))
	for _, bit := range initialDiskState {
		copyOriginal = append(copyOriginal, bit)
	}

	slices.Reverse(copyOriginal)

	convertedInitialState := make([]rune, 0, len(initialDiskState))
	for _, bit := range copyOriginal {
		var transformedBit rune
		if bit == '0' {
			transformedBit = '1'
		} else if bit == '1' {
			transformedBit = '0'
		}

		convertedInitialState = append(convertedInitialState, transformedBit)
	}

	return initialDiskState + "0" + string(convertedInitialState)
}

func fillDisk(initialState string, requiredLength int) string {
	dragonCurve := generateDragonCurve(initialState)
	if len(dragonCurve) >= requiredLength {
		return dragonCurve[:requiredLength]
	}

	return fillDisk(dragonCurve, requiredLength)
}

func generateChecksumCurrentState(state string) string {
	output := []rune{}
	for i := 1; i < len(state); i += 2 {
		nextChar := rune('1')
		if state[i] != state[i-1] {
			nextChar = rune('0')
		}
		output = append(output, nextChar)
	}
	return string(output)
}

func generateChecksum(state string) string {
	currentChecksum := state
	for len(currentChecksum)%2 == 0 {
		currentChecksum = generateChecksumCurrentState(currentChecksum)
	}

	return currentChecksum
}

func part1(initialState string) {
	diskContent := fillDisk(initialState, diskSizePartOne)
	fmt.Printf("Part 1: %s\n", generateChecksum(diskContent))
}

func part2(initialState string) {
	diskContent := fillDisk(initialState, diskSizePartTwo)
	fmt.Printf("Part 2: %s\n", generateChecksum(diskContent))
}

func main() {

	initialDiskState := "1"

	part1(initialDiskState)
	part2(initialDiskState)
}
