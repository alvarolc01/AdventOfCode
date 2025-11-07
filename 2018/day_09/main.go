package main

import (
	"container/ring"
	"fmt"
	"slices"
)

const (
	players int = 10
	marbles int = 1618
)

func gameMaxScore(numMarbles int) int {
	scores := make([]int, players)
	curr := ring.New(1)
	curr.Value = 0

	for i := 1; i <= numMarbles; i++ {
		if i%23 == 0 {
			curr = curr.Move(-7)
			player := (i - 1) % players
			scores[player] += i + curr.Value.(int)

			next := curr.Next()
			curr.Prev().Link(next)
			curr = next
		} else {
			newMarble := ring.New(1)
			newMarble.Value = i

			curr = curr.Move(1)
			curr.Link(newMarble)
			curr = newMarble
		}
	}
	return slices.Max(scores)
}

func part1() {
	fmt.Printf("Part 1: %d\n", gameMaxScore(marbles))
}

func part2() {
	fmt.Printf("Part 2: %d\n", gameMaxScore(marbles*100))
}

func main() {
	part1()
	part2()
}
