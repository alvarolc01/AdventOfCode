package main

import (
	"container/ring"
	"fmt"
)

const (
	maxValuePartOne int = 2017
	maxValuePartTwo int = 50000000
	steps           int = 3
)

func part1() {
	curr := ring.New(1)
	curr.Value = 0

	for i := 1; i <= maxValuePartOne; i++ {
		newMarble := ring.New(1)
		newMarble.Value = i

		curr = curr.Move(steps)
		curr.Link(newMarble)
		curr = newMarble
	}

	fmt.Printf("Part 1: %d\n", curr.Next().Value)
}

func part2() {
	pos := 0
	valAfterZero := 0

	for i := 1; i <= maxValuePartTwo; i++ {
		pos = (pos + steps) % i
		if pos == 0 {
			valAfterZero = i
		}
		pos++
	}

	fmt.Printf("Part 2: %d\n", valAfterZero)
}

func main() {
	part1()
	part2()
}
