package main

import (
	"fmt"
	"strconv"
)

const (
	rangeStart int = 240920
	rangeEnd   int = 789857
)

func countValidPasswords(isPasswordValid func(int) bool) int {
	validPasswords := 0

	for i := rangeStart; i <= rangeEnd; i++ {
		if isPasswordValid(i) {
			validPasswords++
		}
	}

	return validPasswords
}

func areDigitsIncreasing(num string) bool {
	for i := 1; i < len(num); i++ {
		if num[i-1] > num[i] {
			return false
		}
	}

	return true
}

func hasRepeatedDigits(num string) bool {
	for i := 1; i < len(num); i++ {
		if num[i-1] == num[i] {
			return true
		}
	}

	return false
}

func hasExactlyTwoDigits(num string) bool {
	pos := 0

	for pos < len(num) {
		count := 1
		for pos+count < len(num) && num[pos] == num[pos+count] {
			count++
		}

		if count == 2 {
			return true
		}

		pos += count
	}

	return false
}

func part1() {
	isPasswordValid := func(n int) bool {
		convNum := strconv.Itoa(n)
		return hasRepeatedDigits(convNum) && areDigitsIncreasing(convNum)
	}

	validPasswords := countValidPasswords(isPasswordValid)

	fmt.Printf("Part 1: %d\n", validPasswords)
}

func part2() {
	isPasswordValid := func(n int) bool {
		convNum := strconv.Itoa(n)
		return hasRepeatedDigits(convNum) && areDigitsIncreasing(convNum) && hasExactlyTwoDigits(convNum)
	}

	validPasswords := countValidPasswords(isPasswordValid)

	fmt.Printf("Part 2: %d\n", validPasswords)
}

func main() {
	part1()
	part2()
}
