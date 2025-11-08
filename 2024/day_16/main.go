package main

import (
	"container/heap"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	north = iota
	east
	south
	west
)

var directions = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

type Node struct {
	totalCost int
	x, y      int
	direction int
	index     int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].totalCost < pq[j].totalCost }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i]; pq[i].index, pq[j].index = i, j }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Node)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[:n-1]
	return node
}

func findPosition(plot []string, marker rune) (int, int) {
	for x, row := range plot {
		for y, cell := range row {
			if cell == marker {
				return x, y
			}
		}
	}
	return -1, -1
}

func costTurn(ini, end int) int {
	if ini == end {
		return 1
	}
	turns := (end - ini) % 4
	if turns < 0 {
		turns += 4
	}
	if turns > 2 {
		turns = 4 - turns
	}
	return 1000 * turns
}

func bfs(plot []string) {
	iniX, iniY := findPosition(plot, 'S')
	rows, cols := len(plot), len(plot[0])

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Node{
		totalCost: 0,
		x:         iniX,
		y:         iniY,
		direction: east,
	})

	bestCosts := make(map[[3]int]int)
	bestPathDistance := -1
	finalStates := make(map[[3]int]bool)
	previous := make(map[[3]int][][3]int)

	for pq.Len() > 0 {
		top := heap.Pop(pq).(*Node)

		if plot[top.x][top.y] == '#' {
			continue
		}

		if bestPathDistance != -1 && top.totalCost > bestPathDistance {
			break
		}

		state := [3]int{top.x, top.y, top.direction}
		if prevCost, ok := bestCosts[state]; ok && prevCost < top.totalCost {
			continue
		}
		bestCosts[state] = top.totalCost

		if plot[top.x][top.y] == 'E' {
			if bestPathDistance == -1 {
				bestPathDistance = top.totalCost
			}
			if top.totalCost == bestPathDistance {
				finalStates[state] = true
			}
			continue
		}

		for idxDir, dir := range directions {
			nx, ny := top.x+dir[0], top.y+dir[1]
			if idxDir != top.direction {
				nx, ny = top.x, top.y
			}
			if nx < 0 || ny < 0 || nx >= rows || ny >= cols || plot[nx][ny] == '#' {
				continue
			}

			nextCost := top.totalCost + costTurn(top.direction, idxDir)
			nextState := [3]int{nx, ny, idxDir}

			prevCost, exists := bestCosts[nextState]
			if exists && prevCost < nextCost {
				continue
			}

			if !exists || nextCost < prevCost {
				bestCosts[nextState] = nextCost
				previous[nextState] = [][3]int{}
			}

			previous[nextState] = append(previous[nextState], state)

			heap.Push(pq, &Node{
				totalCost: nextCost,
				x:         nx,
				y:         ny,
				direction: idxDir,
			})

		}

	}

	positionsBestPaths := make(map[[2]int]bool)
	visited := make(map[[3]int]bool)

	queue := make([][3]int, 0, len(finalStates))
	for endState := range finalStates {
		queue = append(queue, endState)
		visited[endState] = true
	}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		positionsBestPaths[[2]int{state[0], state[1]}] = true

		for _, parent := range previous[state] {
			if !visited[parent] {
				visited[parent] = true
				queue = append(queue, parent)
			}
		}
	}

	fmt.Printf("Part 1: %d\n", bestPathDistance)
	fmt.Printf("Part 2: %d\n", len(positionsBestPaths))
}

func main() {
	fileName := flag.String("file", "", "Path to the file")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("file name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	content, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("error reading input file:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")

	bfs(lines)
}
