package main

import (
	"fmt"
	"strings"
)

const (
	requiredReceipes int    = 864801
	target           string = "864801"
)

func part1() {
	scores := []int{3, 7}
	firstElf, secondElf := 0, 1

	for len(scores) < requiredReceipes+10 {
		nextVal := scores[firstElf] + scores[secondElf]
		if nextVal >= 10 {
			scores = append(scores, nextVal/10)
		}
		scores = append(scores, nextVal%10)

		firstElf = (firstElf + 1 + scores[firstElf]) % len(scores)
		secondElf = (secondElf + 1 + scores[secondElf]) % len(scores)
	}

	result := scores[requiredReceipes : requiredReceipes+10]
	fmt.Print("Part 1: ")
	for _, n := range result {
		fmt.Print(n)
	}
	fmt.Println()

}

func part2() {
	scores := []rune{'3', '7'}

	firstElf := 0
	secondElf := 1

	for {
		nextVal := int(scores[firstElf]-'0') + int(scores[secondElf]-'0')
		if nextVal >= 10 {
			scores = append(scores, rune('0'+(nextVal/10)))
		}
		scores = append(scores, rune('0'+(nextVal%10)))

		firstElf = (firstElf + 1 + int(scores[firstElf]-'0')) % len(scores)
		secondElf = (secondElf + 1 + int(scores[secondElf]-'0')) % len(scores)

		start := max(0, len(scores)-len(target)-1)
		if idx := strings.Index(string(scores[start:]), target); idx != -1 {
			fmt.Println("Part 2:", start+idx)
			break
		}
	}
}

func main() {
	part1()
	part2()
}
