package main

import (
	"container/list"
	"fmt"
)

const (
	numElfs int = 5
)

func part1() {
	startingElfs := make([]int, 0, numElfs)
	for i := 1; i <= numElfs; i++ {
		startingElfs = append(startingElfs, i)
	}

	for len(startingElfs) > 1 {
		startingElfs = append(startingElfs[2:], startingElfs[0])
	}

	fmt.Printf("Part 1: %d\n", startingElfs[0])
}

func part2() {
	left := list.New()
	right := list.New()

	for i := 1; i <= numElfs; i++ {
		if i <= numElfs/2 {
			left.PushBack(i)
		} else {
			right.PushBack(i)
		}
	}

	for left.Len()+right.Len() > 1 {
		if left.Len() > right.Len() {
			left.Remove(left.Back())
		} else {
			right.Remove(right.Front())
		}

		right.PushBack(left.Remove(left.Front()))

		if right.Len() > left.Len() {
			left.PushBack(right.Remove(right.Front()))
		}
	}

	if left.Len() > 0 {
		fmt.Println("Part 2:", left.Front().Value)
	} else {
		fmt.Println("Part 2:", right.Front().Value)
	}
}

func main() {
	part1()
	part2()
}
