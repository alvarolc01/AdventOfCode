package main

import (
	"fmt"
	"math/bits"
)

const (
	TargetX         = 7
	TargetY         = 4
	StartingX       = 1
	StartingY       = 1
	FavouriteNumber = 10
)

type Point struct {
	x, y int
}

type Node struct {
	pos  Point
	dist int
}

func isWall(p Point) bool {
	x, y := p.x, p.y
	num := x*x + 3*x + 2*x*y + y + y*y
	num += FavouriteNumber

	setBits := bits.OnesCount(uint(num))
	return setBits%2 == 1
}

func getSurroundingPoints(n Node) []Node {
	directions := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	nodes := make([]Node, 0, len(directions))

	for _, d := range directions {
		nodes = append(nodes, Node{
			pos:  Point{n.pos.x + d.x, n.pos.y + d.y},
			dist: n.dist + 1,
		})
	}
	return nodes
}

func getVisitedPoints(stopCondition func(Node) bool) map[Point]int {
	queue := []Node{{Point{StartingX, StartingY}, 0}}
	visited := make(map[Point]int)
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if top.pos.x < 0 || top.pos.y < 0 || isWall(top.pos) {
			continue
		}
		if _, ok := visited[top.pos]; ok {
			continue
		}
		visited[top.pos] = top.dist

		if stopCondition(top) {
			break
		}

		surroundingPoints := getSurroundingPoints(top)
		queue = append(queue, surroundingPoints...)
	}
	return visited
}

func part1() {
	visited := getVisitedPoints(func(n Node) bool {
		return n.pos.x == TargetX && n.pos.y == TargetY
	})
	fmt.Printf("Part 1: %d\n", visited[Point{TargetX, TargetY}])

}

func part2() {
	visited := getVisitedPoints(func(n Node) bool {
		return n.dist > 50
	})
	fmt.Printf("Part 2: %d\n", len(visited)-1)
}

func main() {
	part1()
	part2()
}
